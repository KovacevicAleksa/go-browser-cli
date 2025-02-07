package io

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ReadFile reads a file from the specified path within the "user_files" directory.
// It first constructs the full file path, checks if the file exists, and then reads and prints its contents.
func ReadFile(name string, path string) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file path
	filePath := filepath.Join(folderName, path, name)

	if PathExists(filePath) {
		log.Printf("WARN: file dont exist %s", filePath)
	}

	fmt.Printf("Reading file: %s\n", filePath)

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("ERROR: Error reading file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("File name: %s\n", filePath)
	fmt.Printf("File size: %d bytes\n", len(data))
	fmt.Printf("File content:\n%s\n", string(data))
}
