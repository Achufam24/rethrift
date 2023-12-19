package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Define a simple struct to represent data
type Message struct {
	Text      string `json:"text"`
	FetchTime string
}

type Vertex struct {
	X int
	Y int
}

func logRequestMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Log the request information
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler in the chain
		handler.ServeHTTP(w, r)

		// Log the request duration
		duration := time.Since(startTime)
		log.Printf("Completed %s in %v", r.URL.Path, duration)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	message := fmt.Sprintf("Hello, %s!", name)
	currentTime := time.Now()
	response := Message{Text: message, FetchTime: currentTime.String()}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func main() {
	// Define a route handler function
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// All Apis
	http.Handle("/", logRequestMiddleware(http.HandlerFunc(handler)))
	http.Handle("api/hello", logRequestMiddleware(http.HandlerFunc(helloHandler)))

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on :8080...")

	// Start the server with the loggedHandler
	log.Fatal(http.ListenAndServe(":8080", nil))
}
