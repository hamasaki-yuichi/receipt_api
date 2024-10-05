package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"context"
	"io"
	"log"
	"os"

	// "io"
	// "os"
	// "bufio"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func getUploadFiles() []string {
	credentialFilePath := "/app/key.json"

	// Create client.
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// Get filename from Storage.
	bucket := "receipt-bucket-11"

	it := client.Bucket(bucket).Objects(ctx, nil)
	var names []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
	}
	return names
}

func uploadImageToCloudStorage(f *os.File, name string) string {
	credentialFilePath := "/app/key.json"

	// Create client.
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// Copy to Storage
	bucket := "receipt-bucket-11"
	dir := "temp/"
	object := dir + name

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Errorf("io.Copy:%v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("wc.Close:%v", err)
	}

	fmt.Printf("Blob %v uploaded \n", object)

	log.Println("done")

	return bucket + ":" + object
}

func uploadImageToLocal(w http.ResponseWriter, r *http.Request) string {
	file, fileHeader, err := r.FormFile("file")

	// upload image.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return ""
	}
	defer file.Close()

	err = os.MkdirAll("./uploadimages", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	imageName := time.Now().UnixNano()
	ext := filepath.Ext(fileHeader.Filename)
	dst, err3 := os.Create(fmt.Sprintf("./uploadimages/%d%s", imageName, ext))
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return ""
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	return fmt.Sprintf("%d%s", imageName, ext)
}

// func convertToBase64() string {
//     img, _, _ := image.Decode(file)
//     buffer := new(bytes.Buffer)

//     log.Print("before debug!!!.")
//     log.Printf("img : %#v, %T",img,img)
//     log.Printf("file :  %T",file)
//     var opt jpeg.Options
//     opt.Quality = 80
//     jpeg.Encode(buffer, img, &opt)
//     log.Print("debug!!!.")

//     if err2 := jpeg.Encode(buffer, img, nil); err2 != nil {
//         log.Fatalln("Unable to encode image.")
//     }
//     return base64.StdEncoding.EncodeToString(buffer.Bytes())
// }

// func uploadWithResize(){
//     file, _, err := r.FormFile("file")
//     if err != nil {
//         log.Println(err)
//         http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//         return
//     }

//     img, _, err := image.Decode(file)
//     if err != nil {
//         log.Println(err)
//         http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
//         return
//     }

//     // m := resize.Resize(1000, 0, img, resize.Lanczos3)

//     out, err := os.Create("test_resized.jpg")
//     if err != nil {
//         log.Println(err)
//         http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//         return
//     }
//     defer out.Close()

//     // Encode into jpeg http://blog.golang.org/go-image-package
//     err = jpeg.Encode(out, img, nil)
//     if err != nil {
//         log.Println(err)
//         http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//         return
//     }

//     fmt.Printf("out: %#v",out)
//     return out
// }

// func convertAllImages() string {
// 	var decodeAllImages []image.Image
// 	for _, imageName := range allImageNames {
// 		file, _ := os.Open("assets/" + imageName)
// 		defer file.Close()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		decodeImage, _, err := image.Decode(file)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		decodeAllImages = append(decodeAllImages, decodeImage)
// 	}

// 	var encordImages []string
// 	for _, decodeImage := range decodeAllImages {
// 		buffer := new(bytes.Buffer)
// 		if err := jpeg.Encode(buffer, decodeImage, nil); err != nil {
// 			log.Fatalln("Unable to encode image.")
// 		}
// 		str := base64.StdEncoding.EncodeToString(buffer.Bytes())
// 		encordImages = append(encordImages, str)
// 	}
// }
