package Boot

import (
	"fmt"
	"os"
	"path/filepath"
)

// setupFolder creates a folder if it doesn't exist
func setupFolder(folderName string) error {
	_, err := os.Stat(folderName)
	if err == nil {
		// Folder exists, no need to create it
		return nil
	} else if os.IsNotExist(err) {
		// Folder doesn't exist, create it
		err = os.MkdirAll(folderName, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("INFO: Folder %s created successfully\n", folderName)
		return nil
	} else {
		return err
	}
}

// setupFile creates a file if it doesn't exist
func setupFile(filePath string) error {
	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	if err := setupFolder(dir); err != nil {
		return err
	}

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err == nil {
		// File exists, no need to create it
		return nil
	} else if os.IsNotExist(err) {
		// File doesn't exist, create it
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Printf("INFO: File %s created successfully\n", filePath)
		return nil
	} else {
		return err
	}
}
