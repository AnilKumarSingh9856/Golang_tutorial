package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Server started on port 8000...")

	// 1. GET Handler - Echoes Query Parameters
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handing GET Request")
		
		// Parse query params (e.g., ?name=anil&id=123)
		query := r.URL.Query()
		
		// Create a message string to send back
		message := fmt.Sprintf("Server received GET params: %v", query)
		
		fmt.Println(message)
		fmt.Fprint(w, message)
	})

	// 2. POST Handler - Echoes Body
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handing POST Request")
		echoBody(w, r)
	})

	// postform 
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handing POST Request")
		echoBody(w, r)
	})

	// 3. UPDATE (PUT) Handler - Echoes Body
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handing UPDATE Request")
		echoBody(w, r)
	})

	// 4. DELETE Handler - Echoes Body
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handing DELETE Request")
		echoBody(w, r)
	})

	http.ListenAndServe(":8000", nil)
}

// Helper function to read body and write it back
func echoBody(w http.ResponseWriter, r *http.Request) {
	// Read the body sent by client
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Print to server console
	fmt.Printf("Server received payload: %s\n", string(body))

	// Write back to client
	w.Write(body)
}