package storage

import "sync"

type Storage struct {
	sync.RWMutex
}

type Entry struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	CreatedAt   int64  `json:"createdAt"`
	ModifiedAt  int64  `json:"modifiedAt"`
	Filename    string `json:"filename"`
	Extension   string `json:"extension"`
	Path        string
	MimeType    string
}

type Entries struct {
	Entries []*Entry `json:"entries"`
	Authors []string `json:"authors"`
}
