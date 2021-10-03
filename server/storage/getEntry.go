package storage

import "os"

func (s *Storage) GetEntry(ID string) (*Entry, error) {
	s.RLock()
	defer s.RUnlock()
	entries, err := s.readEntries()
	if err != nil {
		return nil, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for _, entry := range entries.Entries {
		if entry.ID == ID {
			entry.Path = wd + PATH_SEPERATOR + entry.Path
			return entry, nil
		}
	}

	return nil, nil
}
