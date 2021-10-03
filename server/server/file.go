package server

import (
	"fmt"
	"net/http"
)

func (s *Server) File(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		writeError(w, fmt.Errorf("methode non authorisé"), http.StatusMethodNotAllowed)
		return
	}

	ID := req.URL.Query().Get("id")
	entry, err := s.Storage.GetEntry(ID)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	if entry == nil {
		writeError(w, fmt.Errorf("not found"), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+entry.Name+"."+entry.Extension)
	w.Header().Set("Content-Type", entry.MimeType)
	http.ServeFile(w, req, entry.Path)
}
