package storage

import (
	"os"
	"time"

	"github.com/fouad1998/min-system-information/env"
)

func (s *Storage) AddEntry(name, description, owner, extension, mime string, fileContent []byte) (*Entry, error) {
	storagefilepath, err := s.storagefilepath()
	if err != nil {
		return nil, err
	}

	entries, err := s.ReadEntries()
	if err != nil {
		return nil, err
	}

	entryID := s.generateID(entries, func(entry *Entry) string {
		return entry.ID
	})
	filename := s.generateID(entries, func(entry *Entry) string {
		return entry.Filename
	})
	documentPath, err := s.filePath(filename)
	if err != nil {
		return nil, err
	}

	entry := Entry{
		ID:          entryID,
		Name:        name,
		Description: description,
		Owner:       owner,
		Extension:   extension,
		CreatedAt:   time.Now().UnixMilli(),
		ModifiedAt:  time.Now().UnixMilli(),
		MimeType:    mime,
		Path:        env.Setting.StorageDirectory + PATH_SEPERATOR + filename,
		Filename:    filename,
	}
	s.Lock()
	defer s.Unlock()
	documentFile, err := os.OpenFile(documentPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}

	defer documentFile.Close()
	_, err = documentFile.Write(fileContent)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(storagefilepath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	_, err = file.WriteString(s.entryToLine(&entry))
	if err != nil {
		return nil, err
	}

	return &entry, nil
}
