package IO

import (
	"fmt"
	"os"
)

func CreateFile(name string, text string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", name, err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file %s: %v\n", name, cerr)
		}
	}()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", name, err)
		return
	}

	fmt.Printf("File created successfully: %s\n", file.Name())
}
