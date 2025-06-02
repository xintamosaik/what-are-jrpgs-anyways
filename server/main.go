package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from the app!")
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the JRPGs app!")
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
// ...existing code...