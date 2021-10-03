package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) RemoveEntry(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		writeError(w, fmt.Errorf("Methode non autorisé"), http.StatusMethodNotAllowed)
		return
	}

	url := req.URL.Query()
	var params struct {
		ID string `json:"id"`
	}
	params.ID = url.Get("id")
	err := s.Storage.RemoveEntry(params.ID)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Message{
		Message: "Element supprimer avec success",
	})
}
