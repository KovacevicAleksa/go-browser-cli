package IO

import (
	"fmt"
	"go-browser/utils"
	"log"
	"os"
)

// CreateFolder ensures that a folder exists at the specified path and optionally creates it if it doesn't.
func CreateFolder(folderName string) error {
	if folderName == "" {
		return fmt.Errorf("folder name cannot be empty")
	}

	// Check if the folder exists
	if !PathExists(folderName) {
		fmt.Printf("The folder '%s' does not exist.\n", folderName)
		if utils.UserWriteBool("Create folder? (true/false): ") {
			// Create the folder
			if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
				return fmt.Errorf("error creating folder '%s': %w", folderName, err)
			}
			log.Printf("Folder created successfully: %s", folderName)
		} else {
			fmt.Println("Folder creation canceled by the user.")
			return fmt.Errorf("folder creation was canceled")
		}
	} else {
		fmt.Printf("The folder '%s' already exists.\n", folderName)
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
