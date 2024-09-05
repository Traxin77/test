package main

import (
	"fmt"
	"net/http"
)

// Simple handler function
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler) 
	fmt.Println("program works ")
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
