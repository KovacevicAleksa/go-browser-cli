package IO

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreateFile creates a file with the specified name and writes the text to it in a user directory.
func CreateFile(name string, text string, path string) error {
	// Define the folder name
	folderName := "user_files/."
	if path != "" {
		folderName = filepath.Join(folderName, path)
	}
	// Ensure the folder exists with appropriate permissions
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating folder %s: %w", folderName, err)
	}

	// Build the full file path
	filePath := filepath.Join(folderName, name)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("error closing file %s: %v", filePath, cerr)
		}
	}()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, err)
	}

	log.Printf("File created successfully: %s", file.Name())
	return nil
}
