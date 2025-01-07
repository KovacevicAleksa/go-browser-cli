package IO

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteFile(name string) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file path
	filePath := filepath.Join(folderName, name)

	// Attempt to delete the file
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Successfully removed file: %s\n", filePath)
}
