package IO

import (
	"fmt"
	"log"
	"os"
)

// CreateFolder ensures that a folder exists at the specified path and creates it if it doesn't
func CreateFolder(folderName string) error {
	if folderName == "" {
		return fmt.Errorf("folder name cannot be empty")
	}

	// Check if the folder exists
	if !PathExists(folderName) {
		// Create the folder without asking for confirmation
		if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
			return fmt.Errorf("error creating folder '%s': %w", folderName, err)
		}
		log.Printf("Folder created successfully: %s", folderName)
	}

	return nil
}

// PathExists checks if a Path exists
func PathExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
