package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func (s *Server) AddEntry(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	req.ParseMultipartForm(32 << 20) // limit your max input length!
	file, header, err := req.FormFile("file")
	if err != nil {
		writeError(w, fmt.Errorf("pas du fichier attaché à la requête"), http.StatusBadRequest)
		return
	}

	reg, err := regexp.Compile(`.+\.(.+)$`)
	if err != nil {
		writeError(w, fmt.Errorf("une erreur survenu au niveau de system, svp reessayer de nouveau"), http.StatusInternalServerError)
		return
	}

	name := req.Form.Get("name")
	description := req.Form.Get("description")
	owner := req.Form.Get("owner")
	if name == "" || description == "" || owner == "" {
		writeError(w, fmt.Errorf("des données manquantes"), http.StatusBadRequest)
		return
	}

	extension := reg.ReplaceAllString(header.Filename, "$1")
	mime := header.Header.Get("Content-Type")
	io.Copy(&buf, file)
	entry, err := s.Storage.AddEntry(name, description, owner, extension, mime, buf.Bytes())
	if err != nil {
		writeError(w, fmt.Errorf("une erreur survenu au niveau de system, svp reessayer de nouveau"), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entry)
}
