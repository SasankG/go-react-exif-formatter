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

	log.Println("Connected -> api")

	r.ParseMultipartForm(32 << 20)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		log.Println(err)
	}

	file, multipartFileHeader, err := r.FormFile("myImage")
	if err != nil {
		log.Print(err)
		log.Println("failed to retrieve image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	}

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

	util.Transform(imageOutput.Name())

	defer imageOutput.Close()
	defer file.Close()

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
		}
	}

	finalImage, err := os.Open(getImg)
	if err != nil {
		log.Fatal(err)
	}

	defer finalImage.Close()

	imgInfo, _ := finalImage.Stat()
	var finalImageSize = imgInfo.Size()
	finalBuf := make([]byte, finalImageSize)

	fReader := bufio.NewReader(finalImage)
	fReader.Read(finalBuf)

	imgBase64Str := base64.StdEncoding.EncodeToString(finalBuf)

	fmt.Fprintf(w, imgBase64Str)

	os.Remove(imageOutput.Name())

	os.RemoveAll(getImg)

	return

}

func main() {

	r := mux.NewRouter()

	buildHandler := http.FileServer(http.Dir("./client/build"))
	r.PathPrefix("/").Handler(buildHandler).Methods("GET")

	r.HandleFunc("/api", api).Methods("GET", "POST", "OPTIONS")

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on PORT" + port)
	log.Fatal(srv.ListenAndServe())

}
