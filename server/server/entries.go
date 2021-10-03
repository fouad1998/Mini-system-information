package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) Entries(w http.ResponseWriter, req *http.Request) {
	entries, err := s.Storage.ReadEntries()
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entries)
}
