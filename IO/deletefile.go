package IO

import (
	"log"
	"os"
	"path/filepath"
)

// DeleteFile removes the file from the user_files folder and logs the operation.
func DeleteFile(name string) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file path
	filePath := filepath.Join(folderName, name)

	// Attempt to delete the file
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("ERROR: Error deleting file %s: %v", filePath, err)
		return
	}

	log.Printf("INFO: Successfully removed file: %s", filePath)
}
