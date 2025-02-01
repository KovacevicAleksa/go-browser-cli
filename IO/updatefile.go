package IO

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func UpdateFile(name string, text string, overwrite bool) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file path
	filePath := filepath.Join(folderName, name)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("WARN: File %s does not exist.\n", filePath)
		return
	}

	// Determine the file opening mode based on the overwrite parameter
	fileMode := os.O_WRONLY
	if overwrite {
		fileMode |= os.O_TRUNC // Truncate the file if overwrite is true
	} else {
		fileMode |= os.O_APPEND // Append to the file if overwrite is false
	}

	// Open the file with the appropriate mode
	file, err := os.OpenFile(filePath, fileMode, 0644)
	if err != nil {
		log.Printf("ERROR: opening file %s: %v\n", filePath, err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("ERROR: closing file %s: %v\n", filePath, cerr)
		}
	}()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		log.Printf("ERROR: writing to file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("File updated successfully: %s\n", file.Name())
}
