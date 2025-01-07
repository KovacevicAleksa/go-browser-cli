package IO

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(name string, text string) {
	// Define the folder name
	folderName := "user_files"

	// Ensure the folder exists
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating folder %s: %v\n", folderName, err)
		return
	}

	// Build the full file path
	filePath := filepath.Join(folderName, name)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file %s: %v\n", filePath, cerr)
		}
	}()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("File created successfully: %s\n", file.Name())
}
