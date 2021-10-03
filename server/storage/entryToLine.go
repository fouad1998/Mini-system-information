package storage

import (
	"strconv"

	"github.com/fouad1998/min-system-information/env"
)

func (s *Storage) entryToLine(entry *Entry) string {
	se := env.Setting.Seperator
	return entry.ID + se + entry.Name + se + entry.Description + se + entry.Owner + se + strconv.Itoa(int(entry.CreatedAt)) + se + strconv.Itoa(int(entry.ModifiedAt)) + se + entry.Filename + se + entry.Extension + se + entry.Path + se + entry.MimeType + "\n"
}
