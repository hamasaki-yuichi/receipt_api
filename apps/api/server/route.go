package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/show", 301)
}

func index(w http.ResponseWriter, r *http.Request) {
	names := getUploadFiles()

	// render
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, names)
}

func show(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/show.html"))
	data := []Data{}

	tmpl.Execute(w, data)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("template/upload.html"))
		data := []Data{}

		tmpl.Execute(w, data)
	} else if r.Method == "POST" {
		file, fileHeader, err := r.FormFile("file")

		filename := fileHeader.Filename

		tempPath := "images/" + filename
		tmpImage, err := os.Create("images/" + filename)
		if err != nil {
			log.Println(err, "error occurred.")
			return
		}
		defer tmpImage.Close()

		defer file.Close()
		var length int64 = 0
		var img []byte = make([]byte, 1024)
		for {
			n, e := file.Read(img)
			if n == 0 {
				log.Println(e)
				break
			}
			if e != nil {
				log.Println(e)
				return
			}
			tmpImage.WriteAt(img, length)
			length = int64(n) + length
		}

		log.Printf("%#v\n", file)

		// path := uploadImageToLocal(w, r)
		path := uploadImageToCloudStorage(tmpImage, filename)
		err = os.Remove(tempPath)
		if err != nil {
			log.Println(err)
			return
		}

		// render
		tmpl := template.Must(template.ParseFiles("template/upload_success.html"))
		tmpl.Execute(w, path)
	}
}
