package server

import "github.com/fouad1998/min-system-information/storage"

type Server struct {
	Storage storage.Storage
}

type Message struct {
	Message string `json:"message"`
}
