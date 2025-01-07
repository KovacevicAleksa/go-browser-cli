package IO

import (
	"fmt"
	"os"
	"path/filepath"
)

func RenameFile(oldName string, newName string) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file paths
	oldFilePath := filepath.Join(folderName, oldName)
	newFilePath := filepath.Join(folderName, newName)

	// Check if the old file exists
	if _, err := os.Stat(oldFilePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", oldFilePath)
		return
	}

	// Rename the file
	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		fmt.Printf("Error renaming file from %s to %s: %v\n", oldFilePath, newFilePath, err)
		return
	}

	fmt.Printf("File renamed successfully from %s to %s\n", oldFilePath, newFilePath)
}
