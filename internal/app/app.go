package app

import (
    "net/http"
    "encoding/json"
)

// HelloResponse represents the response structure for the hello message.
type HelloResponse struct {
    Message string `json:"message"`
}

// HelloHandler handles the "hello" message request.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := HelloResponse{Message: "Hello from the app!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}