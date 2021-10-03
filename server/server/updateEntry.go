package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func (s *Server) UpdateEntry(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		writeError(w, fmt.Errorf("Methode non autorisé"), http.StatusMethodNotAllowed)
		return
	}

	var buf bytes.Buffer
	var (
		extension, mime string
	)
	req.ParseMultipartForm(32 << 20) // limit your max input length
	file, header, err := req.FormFile("file")
	if err == nil {
		reg, err := regexp.Compile(`.+\.(.+)$`)
		if err != nil {
			writeError(w, fmt.Errorf("une erreur survenu au niveau de system, svp reessayer de nouveau"), http.StatusInternalServerError)
			return
		}

		extension = reg.ReplaceAllString(header.Filename, "$1")
		mime = header.Header.Get("Content-Type")
		io.Copy(&buf, file)
	}

	ID := req.Form.Get("id")
	name := req.Form.Get("name")
	description := req.Form.Get("description")
	owner := req.Form.Get("owner")
	if name == "" || description == "" || owner == "" || ID == "" {
		writeError(w, fmt.Errorf("des données manquantes"), http.StatusBadRequest)
		return
	}

	entry, err := s.Storage.UpdateEntry(ID, name, description, owner, extension, mime, buf.Bytes())
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entry)
}
