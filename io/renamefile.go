package io

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func RenameFile(oldName string, newName string, path string) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file paths
	oldFilePath := filepath.Join(folderName, path, oldName)
	newFilePath := filepath.Join(folderName, path, newName)

	// Check if the old file exists
	if _, err := os.Stat(oldFilePath); os.IsNotExist(err) {
		log.Printf("WARN: File %s does not exist.\n", oldFilePath)
		return
	}

	// Rename the file
	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		log.Printf("ERROR: Error renaming file from %s to %s: %v\n", oldFilePath, newFilePath, err)
		return
	}

	fmt.Printf("File renamed successfully from %s to %s\n", oldFilePath, newFilePath)
}
