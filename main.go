package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
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

	// obtain image from form data
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

	// logs -------
	fmt.Println("Success, image upload was successful")
	fmt.Println(reflect.TypeOf(file))
	// ------------

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Print(err)
	} else {
		log.Print(reflect.TypeOf(buf))
	}

	// open image and create a file on server
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

	// save file in ./images
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

	// transform image
	util.Transform(imageOutput.Name())

	defer imageOutput.Close()
	defer file.Close()

	// return transformed image to client
	var files []string
	var getImg string

	// walk through transformed image folder
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

	// open image to send
	finalImage, err := os.Open(getImg)
	if err != nil {
		log.Fatal(err)
	}

	defer finalImage.Close()

	// create buffer for image to send
	imgInfo, _ := finalImage.Stat()
	var finalImageSize = imgInfo.Size()
	finalBuf := make([]byte, finalImageSize)

	// read final image into a biffer
	fReader := bufio.NewReader(finalImage)
	fReader.Read(finalBuf)

	// convert to base 64
	imgBase64Str := base64.StdEncoding.EncodeToString(finalBuf)

	// send to client base64
	fmt.Fprintf(w, imgBase64Str)

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
