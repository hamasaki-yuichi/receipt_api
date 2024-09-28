package server

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	log.Print("starting server...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Print("missing port")
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/list", list)

	// fileserver
	fileServer := http.FileServer(http.Dir("./uploadimages"))
	http.Handle("/images/", http.StripPrefix("/images/", fileServer))

	// server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
