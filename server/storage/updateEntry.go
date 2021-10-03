package storage

import (
	"fmt"
	"os"
	"time"
)

func (s *Storage) UpdateEntry(ID, name, description, owner, extension, mime string, fileContent []byte) (*Entry, error) {
	s.Lock()
	defer s.Unlock()
	entries, err := s.readEntries()
	if err != nil {
		return nil, err
	}

	found := false
	filename := ""
	fileIndex := 0
	var createdAt int64 = 0
	path := ""
	extensionq := ""
	mimeq := ""
	for index, entry := range entries.Entries {
		if entry.ID == ID {
			filename = entry.Filename
			fileIndex = index
			extensionq = entry.Extension
			mimeq = entry.MimeType
			path = entry.Path
			found = true
			createdAt = (entry.CreatedAt)
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("non trouve")
	}

	if len(fileContent) != 0 {
		directory, err := s.storageDirectory()
		if err != nil {
			return nil, err
		}

		file, err := os.OpenFile(directory+PATH_SEPERATOR+filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
		if err != nil {
			return nil, err
		}

		file.Write(fileContent)
		file.Close()
		extensionq = extension
		mimeq = mime
	}

	entry := &Entry{
		ID:          ID,
		Name:        name,
		Description: description,
		Owner:       owner,
		Filename:    filename,
		Extension:   extensionq,
		CreatedAt:   createdAt,
		ModifiedAt:  time.Now().UnixMilli(),
		MimeType:    mimeq,
		Path:        path,
	}
	entries.Entries[fileIndex] = entry
	storagefilepath, err := s.storagefilepath()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(storagefilepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	for _, entry := range entries.Entries {
		file.WriteString(s.entryToLine(entry))
	}

	return entry, nil
}
