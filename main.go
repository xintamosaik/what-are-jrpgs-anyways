package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Define a struct for our JSON response
type DirectoryContents struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
}

func create_files_and_folders_json(list []os.DirEntry) (string, error) {
	if list == nil {
		return "{}", nil
	}

	contents := DirectoryContents{
		Folders: make([]string, 0),
		Files:   make([]string, 0),
	}

	for _, file := range list {
		if file.IsDir() {
			contents.Folders = append(contents.Folders, file.Name())
		} else {
			contents.Files = append(contents.Files, file.Name())
		}
	}

	jsonBytes, err := json.Marshal(contents)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	return string(jsonBytes), nil
}

// tile_editor serves the tile editor HTML page.
func list_uploads(w http.ResponseWriter, r *http.Request) {
	// json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	list := list_files_and_folders_of_directory("")
	if list == nil {
		http.Error(w, `{"error": "Error reading uploads directory"}`, http.StatusInternalServerError)
		return
	}

	json, err := create_files_and_folders_json(list)
	if err != nil {
		http.Error(w, `{"error": "Error creating JSON"}`, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, json)
}

// noCache is a middleware that prevents caching of responses.
func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

// // filter_list_of_files_for_extension returns a filtered list of files that match the given extension.
// func filter_list_of_files_for_extension(list []os.DirEntry, extension string) []os.DirEntry {
// 	var filtered []os.DirEntry
// 	for _, file := range list {
// 		if file.IsDir() {
// 			continue // Skip directories
// 		}
// 		if len(file.Name()) >= len(extension) && file.Name()[len(file.Name())-len(extension):] == extension {
// 			filtered = append(filtered, file)
// 		}
// 	}
// 	return filtered
// }

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

	fs := http.StripPrefix("/", http.FileServer(http.Dir("ui/dist")))
	http.Handle("/", noCache(fs))

	fs_uploads := http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/")))
	http.Handle("/uploads/", noCache(fs_uploads)) // Note the trailing slash

	// API route
	http.HandleFunc("/list_uploads", list_uploads)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
