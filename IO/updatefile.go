package IO

import (
	"fmt"
	"os"
)

func UpdateFile(name string, text string) {
	// Check if the file exists before trying to open it
	if _, err := os.Stat(name); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", name)
		return
	}

	// Open the file in read-write mode, create it if it doesn't exist, and append if it does.
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", name, err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file %s: %v\n", name, cerr)
		}
	}()

	// Write the provided text to the file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", name, err)
		return
	}

	fmt.Printf("File updated successfully: %s\n", file.Name())
}
