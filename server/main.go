package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

func main() {
	staticDir := "../ui/dist"

	// Confirm static files exist
	if _, err := os.Stat(staticDir + "/index.html"); os.IsNotExist(err) {
		fmt.Println("ERROR: index.html not found in", staticDir)
	}

	fs := http.StripPrefix("/", http.FileServer(http.Dir(staticDir)))
	http.Handle("/", fs)

	http.HandleFunc("/api/ping", pingHandler)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
