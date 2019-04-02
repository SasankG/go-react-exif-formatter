package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Define the routes
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "hello")
}

func api(w http.ResponseWriter, r *http.Request) {
	// Change route to get mux.Vars and also to accept Multipart content-type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Connected -> api route")

	// The response sending
	sample := "connected, this is a JSON"
	jsonsample, err := json.Marshal(sample)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(jsonsample)
	fmt.Fprintf(w, "connected to api")
}

func main() {

	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc("/", index).Methods("GET")
	// Initially testing GET, Change to POST after
	r.HandleFunc("/api", api).Methods("GET", "POST", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "POST"},
	})

	// Start the server
	fmt.Println("Server started on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
