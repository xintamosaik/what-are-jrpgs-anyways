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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the JRPGs app!")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/ping", pingHandler)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
