package main

import (
	"fmt"
	"net/http"
)

// Define the routes
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing")
	// Run the util functions here
}

// Create webserver
func main() {
	// Routes handles here
	http.HandleFunc("/", index)
	fmt.Println("Server starting on PORT 8080")
	http.ListenAndServe(":8080", nil)
}
