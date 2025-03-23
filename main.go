package main

import (
	"net/http"
)

func main() {
	// Define the folder to serve files from
	fs := http.FileServer(http.Dir("./www"))

	// Handle all requests by serving files from the "www" folder
	http.Handle("/", fs)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
