package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello new World!!")
}

func main() {
	log.Print("starting server...")
	port := os.Getenv("PORT")
    if port == "" {
		log.Print("missing port")
    }

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
