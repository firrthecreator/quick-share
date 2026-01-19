package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Config holds the configuration for the HTTP server.
type Config struct {
	Port      string
	Directory string // Directory to serve or save uploads to
	Upload    bool   // If true, enables upload mode
}

// Start initializes and starts the HTTP server.
func Start(cfg Config) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if cfg.Upload {
			handleUpload(w, r, cfg.Directory)
		} else {
			// Serve files directly for download
			fs := http.FileServer(http.Dir(cfg.Directory))
			fs.ServeHTTP(w, r)
		}
	})

	return http.ListenAndServe(":"+cfg.Port, nil)
}

// handleUpload manages the file upload interface and processing.
func handleUpload(w http.ResponseWriter, r *http.Request, uploadDir string) {
	// 1. GET Request: Show the upload form
	if r.Method == "GET" {
		renderUploadForm(w)
		return
	}

	// 2. POST Request: Process the uploaded file
	if r.Method == "POST" {
		// Limit upload size to 1GB to prevent memory exhaustion
		if err := r.ParseMultipartForm(10 << 30); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create the destination file
		dstPath := filepath.Join(uploadDir, handler.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Error creating destination file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy data from request to file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Received file: %s\n", handler.Filename)

		// Success page
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Upload Successful!</h1><p>File saved as <b>%s</b></p><a href='/'>Upload another</a>", handler.Filename)
	}
}

// renderUploadForm serves a simple HTML page for file uploading.
func renderUploadForm(w http.ResponseWriter) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Quick Share Upload</title>
		<style>
			body { font-family: sans-serif; text-align: center; padding: 50px; }
			form { border: 2px dashed #ccc; padding: 20px; display: inline-block; }
			input { margin: 10px 0; }
			button { background-color: #007bff; color: white; padding: 10px 20px; border: none; cursor: pointer; }
			button:hover { background-color: #0056b3; }
		</style>
	</head>
	<body>
		<h2>Upload File to Host</h2>
		<form action="/" method="post" enctype="multipart/form-data">
			<input type="file" name="file" required /><br><br>
			<button type="submit">Upload</button>
		</form>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write([]byte(html)); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}
