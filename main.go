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

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sasankg/go-exif/util"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "hello")
}

func api(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Connected -> api route")
	fmt.Fprintf(w, "connected to api")

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

	dir := "./images"
	dst, err := os.Create(filepath.Join(dir, filepath.Base(imageOutput.Name())))
	if err != nil {
		log.Fatal(err)
	}

	defer imageOutput.Close()

	imageSave, err := os.Open(imageOutput.Name())
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(dst, imageSave)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Saved")

	util.Transform(imageOutput.Name())

	defer file.Close()

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api", api).Methods("GET", "POST", "OPTIONS")

	// temp cors fix
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})

	fmt.Println("Server started on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
