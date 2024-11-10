package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	imagesDir   = "./images"
	galleryFile = "gallery.html"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, galleryFile)
	})

	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		imageName := r.URL.Path[len("/images/"):]
		imagePath := fmt.Sprintf("%s/%s", imagesDir, imageName)
		file, err := os.Open(imagePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer file.Close()

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(buf[:n])
		}
	})

	http.HandleFunc("/get-images", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		imageFiles := []string{}
		err := filepath.Walk(imagesDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				imageFiles = append(imageFiles, info.Name())
			}
			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(imageFiles)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}