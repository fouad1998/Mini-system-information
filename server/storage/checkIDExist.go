package storage

func (s *Storage) checkIDExists(id string, entries *Entries, grab func(entry *Entry) string) bool {
	for _, entry := range entries.Entries {
		if id == grab(entry) {
			return true
		}
	}

	return false
}
