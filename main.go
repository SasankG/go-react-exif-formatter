package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	// r.FormFile automatically calls parsemultipartform, we are doing it to double check
	r.ParseMultipartForm(32 << 20)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Println(err)
	}
	file, _, err := r.FormFile("myImage")
	if err != nil {
		log.Print(err)
		fmt.Println("failed to retrieve image")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err == nil {
		fmt.Println("Success, image upload was successful")
		fmt.Println(file)
	}

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
