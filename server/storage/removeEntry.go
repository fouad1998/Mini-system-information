package storage

import "os"

func (s *Storage) RemoveEntry(ID string) error {
	s.Lock()
	defer s.Unlock()
	entries, err := s.readEntries()
	if err != nil {
		return err
	}

	found := false
	for index, entry := range entries.Entries {
		if entry.ID == ID {
			directory, err := s.storageDirectory()
			if err != nil {
				return err
			}

			os.Remove(directory + PATH_SEPERATOR + entry.Filename)
			entries.Entries = append(entries.Entries[:index], entries.Entries[index+1:]...)
			found = true
			break
		}
	}

	if found {
		storagefilepath, err := s.storagefilepath()
		if err != nil {
			return err
		}

		file, err := os.OpenFile(storagefilepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return nil
		}

		defer file.Close()
		for _, entry := range entries.Entries {
			file.WriteString(s.entryToLine(entry))
		}
	}

	return nil
}
