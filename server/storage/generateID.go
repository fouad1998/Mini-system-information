package storage

import uuid "github.com/satori/go.uuid"

func (s *Storage) generateID(entries *Entries, grab func(entry *Entry) string) string {
	for {
		id := uuid.NewV4().String()
		exists := s.checkIDExists(id, entries, grab)
		if !exists {
			return id
		}
	}
}
