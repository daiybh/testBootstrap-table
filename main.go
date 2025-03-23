package main

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// 定义数据
	items := []Item{
		{ID: 1, Name: "Item 1", Price: 10.0},
		{ID: 2, Name: "Item 2", Price: 20.0},
		{ID: 3, Name: "Item 3", Price: 30.0},
	}

	// Define the folder to serve files from
	fs := http.FileServer(http.Dir("./www"))

	// Handle all requests by serving files from the "www" folder
	http.Handle("/", fs)
	http.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	})
	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
