package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/sasankg/go-exif/util"
)

func api(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Connected -> api route")

	fmt.Fprintf(w, "hello")

	r.ParseMultipartForm(32 << 20)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Println(err)
	}

	file, multipartFileHeader, err := r.FormFile("myImage")
	if err != nil {
		log.Print(err)
		fmt.Println("failed to retrieve image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Success, image upload was successful")
	fmt.Println(reflect.TypeOf(file))

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Print(err)
	} else {
		log.Print(reflect.TypeOf(buf))
	}

	imageOutput, err := os.Create(multipartFileHeader.Filename)
	if err != nil {
		log.Print(err)
	}

	_, err = imageOutput.Write(buf.Bytes())
	if err != nil {
		log.Print(err)
	} else {
		log.Print(reflect.TypeOf(imageOutput))
		namer := imageOutput.Name()
		log.Print(namer)
	}

	// Save
	// open the file for reading
	imageSave, err := os.Open(imageOutput.Name())
	if err != nil {
		log.Fatal(err)
	}

	dir := "./images"
	dst, err := os.Create(filepath.Join(dir, filepath.Base(util.NameGen(4)+imageSave.Name())))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(dst, imageSave)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File Saved")
	defer imageSave.Close()
	defer dst.Close()

	// transform
	util.Transform(imageOutput.Name())

	defer imageOutput.Close()
	defer file.Close()

	// return image to client
	var files []string
	var getImg string

	root := "./testingsave"
	errs := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if errs != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file) == ".jpg" {
			getImg = file
			log.Println(getImg)
		}
	}

	// change this to the file in the testing save folder
	finalImage, err := os.Open(getImg)
	if err != nil {
		log.Fatal(err)
	}

	defer finalImage.Close()

	// get header parameters and info
	ContentType := http.DetectContentType(buf.Bytes())

	// set headers
	w.Header().Set("Content-Type", ContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+imageOutput.Name())

	// send to the client
	io.Copy(w, finalImage)

	// delete image
	os.Remove(imageOutput.Name())

	// clear testingsave
	os.RemoveAll(getImg)

	return

}

func main() {

	r := mux.NewRouter()

	// handle app
	buildHandler := http.FileServer(http.Dir("./client/build"))
	r.PathPrefix("/").Handler(buildHandler).Methods("GET")

	r.HandleFunc("/api", api).Methods("GET", "POST", "OPTIONS")

	// configure server
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// serve
	fmt.Println("Server started on PORT 8080")
	log.Fatal(srv.ListenAndServe())

}
