package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/joho/godotenv"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}
func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// look for a .env file in the current directory
	// and load it into the environment variables
	env, err := godotenv.Read(".env")
	if err != nil {
		fmt.Println("No .env file found, proceeding without it.")
	}
	if len(env) > 0 {
		fmt.Println("Loaded environment variables from .env file.")
		for key, value := range env {
			fmt.Printf("%s=%s\n", key, value)
		}
	} else {
		fmt.Println("No environment variables loaded from .env file.")
	}
	// If you have a .env file, it will load the environment variables
	IS_DEV := env["DEV"] == "true"
	if IS_DEV {
		fmt.Println("Development mode is ON")

		// Proxy all requests to Vite dev server
		target, _ := url.Parse("http://localhost:5173")
		proxy := httputil.NewSingleHostReverseProxy(target)

		http.Handle("/", proxy)

	} else {
		fs := http.StripPrefix("/", http.FileServer(http.Dir("../ui/dist")))
		http.Handle("/", noCache(fs))

	}
	// Then wrap your file server:

	// API route
	http.HandleFunc("/api/ping", pingHandler)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
