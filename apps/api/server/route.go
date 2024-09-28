package server

import (
	"html/template"
	"net/http"
	"os"
	"log"
)

type Data struct {
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	data := []Data{}

	tmpl.Execute(w, data)
}

func show(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/show.html"))
	data := []Data{}

	tmpl.Execute(w, data)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/show", 301)
	} else if r.Method == "POST" {
		path := uploadImageToLocal(w, r)

		// render
		tmpl := template.Must(template.ParseFiles("template/upload.html"))
		tmpl.Execute(w, path)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Open("uploadimages/")
	defer dir.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	images, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatalln("No files")
	}

	// render
	tmpl := template.Must(template.ParseFiles("template/list.html"))
	tmpl.Execute(w, images)
}
