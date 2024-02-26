package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 1000*1024*1024*1024)

	// Parse the form data, processing multipart form data in chunks of 32MB
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileHeader := r.MultipartForm.File["file"][0]
	filename := fileHeader.Filename

	uploadPath := filepath.Join("./uploads", filepath.Base(filename))
	dst, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file in chunks, avoiding loading the entire file into memory
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
