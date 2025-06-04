package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/a-h/templ"
	"go-web-rpg-app/components" // Import the generated hello component
)

// Define a struct for our JSON response
type DirectoryContents struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
}

func compile_files_and_folders_json(list []os.DirEntry) (string, error) {
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

// list_uploads handles the /list_uploads API endpoint.
// It lists all files and folders in the specified uploads directory.
// The folder parameter can be used to specify a subdirectory within uploads.
// If no folder is specified, it lists the root of the uploads directory.
func list_uploads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	// Get the folder parameter from query string
	folder := r.URL.Query().Get("folder")

	// List files for the requested folder
	list := list_files_and_folders_of_directory(folder)
	if list == nil {
		http.Error(w, `{"error": "Error reading uploads directory"}`, http.StatusInternalServerError)
		return
	}

	json, err := compile_files_and_folders_json(list)
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

// list_files_and_folders_of_directory returns a list of files and folders in a given directory.
// It prevents directory traversal attacks by disallowing double points and other unsafe characters.
// The subdir parameter specifies the subdirectory to list files from. If empty, it lists the root directory.
func list_files_and_folders_of_directory(subdir string) []os.DirEntry {
	basedir := "./uploads/"

	// Basic path sanitization
	if subdir != "" {
		// Remove leading and trailing slashes
		subdir = strings.Trim(subdir, "/")

		// Check for path traversal attempts
		if strings.Contains(subdir, "..") {
			fmt.Println("Error: Path traversal attempt detected")
			return nil
		}

		// Check for invalid characters
		if strings.ContainsAny(subdir, "\\:*?\"<>|") {
			fmt.Println("Error: Invalid characters in path")
			return nil
		}

		basedir = filepath.Join(basedir, subdir)
	}

	// Check if the resulting path is still within uploads directory
	absPath, err := filepath.Abs(basedir)
	if err != nil {
		fmt.Println("Error resolving path:", err)
		return nil
	}

	absUploads, err := filepath.Abs("./uploads")
	if err != nil {
		fmt.Println("Error resolving uploads path:", err)
		return nil
	}

	if !strings.HasPrefix(absPath, absUploads) {
		fmt.Println("Error: Path traversal attempt detected")
		return nil
	}

	files, err := os.ReadDir(basedir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	return files
}


func main() {

	component := components.Hello("John")
	http.Handle("/hello", templ.Handler(component))

	fs := http.StripPrefix("/", http.FileServer(http.Dir("web-vanilla")))
	http.Handle("/", noCache(fs))

	fs_uploads := http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/")))
	http.Handle("/uploads/", noCache(fs_uploads)) // Note the trailing slash

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
