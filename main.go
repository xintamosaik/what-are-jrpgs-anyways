package main

import (
	"fmt"
	"net/http"
	"os"
)

// tile_editor serves the tile editor HTML page.
func tile_editor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store")
	html := `
<div>
<h1>Tile Editor!</h1>
<p>Welcome to the Tile Editor!</p>
</div>
`

	// Serve the HTML content as a response
	fmt.Fprint(w, html)

}

// noCache is a middleware that prevents caching of responses.
func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

// filter_list_of_files_for_extension returns a filtered list of files that match the given extension.
func filter_list_of_files_for_extension(list []os.DirEntry, extension string) []os.DirEntry {
	var filtered []os.DirEntry
	for _, file := range list {
		if file.IsDir() {
			continue // Skip directories
		}
		if len(file.Name()) >= len(extension) && file.Name()[len(file.Name())-len(extension):] == extension {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// list_files_and_folders_of_directory returns a list of files and folders in a given directory.
// It prevents directory traversal attacks by disallowing double points and other unsafe characters.
// The subdir parameter specifies the subdirectory to list files from. If empty, it lists the root directory.
func list_files_and_folders_of_directory(subdir string) []os.DirEntry {
	basedir := "./uploads/"

	// We prevent traversal attacks by searching the whole string for double points
	if len(subdir) > 0 && subdir[0] == '/' {
		fmt.Println("Error: Subdirectory cannot start with a slash.")
		return nil
	}

	if len(subdir) > 0 && subdir[len(subdir)-1] == '/' {
		fmt.Println("Error: Subdirectory cannot end with a slash.")
		return nil
	}
	// Double points are never allowed. We search the string with a loop
	for _, char := range subdir {
		if char == '.' {
			fmt.Println("Error: Subdirectory cannot contain double points.")
			return nil
		}
		if char == '\\' {
			fmt.Println("Error: Subdirectory cannot contain backslashes.")
			return nil
		}

		if char == ':' {
			fmt.Println("Error: Subdirectory cannot contain colons.")
			return nil
		}
		if char == '*' || char == '?' || char == '"' || char == '<' || char == '>' || char == '|' {
			fmt.Println("Error: Subdirectory cannot contain invalid characters.")
			return nil
		}
	}

	if subdir != "" {
		basedir += subdir
	}

	files, err := os.ReadDir(basedir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	return files
}

func main() {
	// try the function to list files and folders in the uploads directory
	rootfiles := list_files_and_folders_of_directory("1_Interiors/16x16")
	for _, file := range rootfiles {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
	fmt.Println("Listing files and folders in the uploads directory:")
	pngs := filter_list_of_files_for_extension(rootfiles, ".png")
	for _, file := range pngs {
		fmt.Println("PNG File:", file.Name())
	}

	fs := http.StripPrefix("/", http.FileServer(http.Dir("./web/")))
	http.Handle("/", noCache(fs))

	fs_uploads := http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/")))
	http.Handle("/uploads/", noCache(fs_uploads)) // Note the trailing slash

	// API route
	http.HandleFunc("/tile_editor", tile_editor)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
