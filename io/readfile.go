package io

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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
