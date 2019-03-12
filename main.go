package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define the routes
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Run the util functions here
}

func api(w http.ResponseWriter, r *http.Request) {
	// Change route to get mux.Vars and also to accept Multipart content-type
	w.Header().Set("Content-Type", "application/json")
}

// Create webserver
func main() {
	// init router
	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api", api).Methods("POST")

	// Start the server
	fmt.Println("Server started on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
