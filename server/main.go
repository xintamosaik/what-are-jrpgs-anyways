package main

import (

	"fmt"
	"net/http"
)

// it will serve html 
func tile_editor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store")
	html := `
<div>
<h1>Tile Editor</h1>
</div>
`

	// Serve the HTML content as a response
	fmt.Fprint(w, html)
	

}
func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {

	fs := http.StripPrefix("/", http.FileServer(http.Dir("./web/")))
	http.Handle("/", noCache(fs))

	// API route
	http.HandleFunc("/tile_editor", tile_editor)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
