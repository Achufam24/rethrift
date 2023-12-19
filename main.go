package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a simple struct to represent data
type Message struct {
	Text string `json:"text"`
}

func main() {
	// Define a route handler function
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		// Create a response struct
		response := Message{Text: "Hello, World!"}

		// Convert the response struct to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the response writer
		w.Write(jsonResponse)
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
