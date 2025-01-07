package IO

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFile(folderName string) {
	// Define the folder path
	baseFolder := "user_files"
	folderPath := filepath.Join(baseFolder, folderName)

	// Open the directory
	dir, err := os.Open(folderPath)
	if err != nil {
		fmt.Printf("Error opening directory %s: %v\n", folderPath, err)
		return
	}
	defer dir.Close()

	// Read the directory entries
	files, err := dir.Readdir(-1) // -1 means read all files and directories
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", folderPath, err)
		return
	}

	// Print the list of files and directories
	fmt.Printf("Contents of directory %s:\n", folderPath)
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("[DIR]  %s\n", file.Name())
		} else {
			fmt.Printf("[FILE] %s\n", file.Name())
		}
	}
}
