package storage

import (
	"bufio"
	"os"
)

func (s *Storage) ReadEntries() (entries *Entries, err error) {
	s.RLock()
	defer s.RUnlock()
	return s.readEntries()
}

func (s *Storage) readEntries() (entries *Entries, err error) {
	entries = &Entries{
		Entries: make([]*Entry, 0),
		Authors: make([]string, 0),
	}
	storagefilepath, err := s.storagefilepath()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(storagefilepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		entry, err := s.extractEntry(line)
		if err != nil {
			return nil, err
		}

		entries.Entries = append(entries.Entries, entry)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	for _, entry := range entries.Entries {
		found := false
		for _, author := range entries.Authors {
			if author == entry.Owner {
				found = true
				break
			}
		}

		if !found {
			entries.Authors = append(entries.Authors, entry.Owner)
		}
	}

	return entries, nil
}
