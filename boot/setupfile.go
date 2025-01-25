package Boot

import (
	"log"
	"os"
)

func setupFolder(folderName string) error {
	// Ensure the folder exists
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating folder %s: %v", folderName, err)
	}
	log.Printf("INFO: Main Folder created successfully: [%s]\n", folderName)
	return nil
}
