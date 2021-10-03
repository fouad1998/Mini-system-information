package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	server := Server{}
	if err := server.Storage.Init(); err != nil {
		log.Fatalln("Faile to start, error: ", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Faile to start, error: ", err)
	}

	go func() {
		log.Println("Serveur prêt à http://localhost:5252")
		log.Fatalln(http.ListenAndServe(":5252", http.FileServer(http.Dir(wd+"\\public"))))
	}()

	router.Handle("/add", cors(server.AddEntry))
	router.Handle("/entries", cors(server.Entries))
	router.Handle("/remove", cors(server.RemoveEntry))
	router.Handle("/update", cors(server.UpdateEntry))
	router.Handle("/file", cors(server.File))
	log.Println("Les services sont prêt à http://localhost:8000")
	http.ListenAndServe(":8000", router)
}
