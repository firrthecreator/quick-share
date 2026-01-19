package server

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestHandleUpload_Get verifies that the GET request returns the upload HTML form.
func TestHandleUpload_Get(t *testing.T) {
	// 1. Create a simulated GET request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// 3. Create a dummy upload directory (not used in GET, but required by signature)
	tempDir := t.TempDir()

	// 4. Call the handler function directly
	handleUpload(rr, req, tempDir)

	// 5. Assertions
	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if body contains specific HTML keywords
	expectedKeyword := "Upload File to Host"
	if !strings.Contains(rr.Body.String(), expectedKeyword) {
		t.Errorf("handler returned unexpected body: does not contain %v",
			expectedKeyword)
	}
}

// TestHandleUpload_Post verifies that the POST request correctly saves the file.
func TestHandleUpload_Post(t *testing.T) {
	// 1. Setup a temporary directory for file upload simulation
	tempDir := t.TempDir()

	// 2. Prepare multipart form data
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a form file field named "file"
	part, err := writer.CreateFormFile("file", "test_document.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Write some content to the "virtual" file
	fileContent := "This is a test content for unit testing."
	_, err = part.Write([]byte(fileContent))
	if err != nil {
		t.Fatal(err)
	}

	// Close the writer to set the boundary
	writer.Close()

	// 3. Create a simulated POST request with the form data
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal(err)
	}
	// Important: Set the Content-Type header with the boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 4. Record the response
	rr := httptest.NewRecorder()

	// 5. Call the handler
	handleUpload(rr, req, tempDir)

	// 6. Assertions
	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify the file was actually created in the temp directory
	expectedPath := filepath.Join(tempDir, "test_document.txt")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Expected file to be created at %s, but it does not exist", expectedPath)
	}

	// Verify file content matches
	savedContent, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(savedContent) != fileContent {
		t.Errorf("File content mismatch: got %s, want %s", string(savedContent), fileContent)
	}
}
