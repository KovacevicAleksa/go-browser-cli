package main

import (
	"go-browser/IO"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	filename := "test_delete.txt"
	content := "This is a test file for deletion."

	// Ensure the folder exists
	os.MkdirAll("user_files", os.ModePerm)

	// Create a file to be deleted
	filePath := filepath.Join("user_files", filename)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file %s: %v", filePath, err)
	}

	// Call the function to delete the file
	IO.DeleteFile(filename)

	// Verify the file was deleted
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Fatalf("File %s was not deleted", filePath)
	}

	// Clean up
	if err := os.Remove("user_files"); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove test folder user_files: %v", err)
	}
}
