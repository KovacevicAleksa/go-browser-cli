package main

import (
	"go-browser/IO"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFile(t *testing.T) {
	filename := "test_create.txt"
	content := "This is a test file."

	err := IO.CreateFile(filename, content, "", true)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Build the expected file path
	expectedFilePath := filepath.Join("user_files", filename)

	// Verify the file was created
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Fatalf("File %s was not created", expectedFilePath)
	}

	// Verify the file content
	data, err := os.ReadFile(expectedFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", expectedFilePath, err)
	}

	if string(data) != content {
		t.Errorf("File content mismatch: expected %q, got %q", content, string(data))
	}

	// Clean up
	if err := os.Remove(expectedFilePath); err != nil {
		t.Fatalf("Failed to remove test file %s: %v", expectedFilePath, err)
	}

	// Remove the folder if empty
	if err := os.Remove("user_files"); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove test folder user_files: %v", err)
	}
}
