package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}


func main() {
	// http.HandleFunc("/", rootHandler) // ui/dist/index.html needs to point to this
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// log the request
		fmt.Println("Received request:", r.Method, r.URL.Path) // we do get here 
		// But we are in /server so we need to adjust the path to serve the index.html file correctly
		// Serve the index.html file from the ui/dist directory
		http.ServeFile(w, r, "../ui/dist/index.html")
	})
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello from the app!"})
	})
	// Register the ping handler
	http.HandleFunc("/api/ping", pingHandler) 

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
