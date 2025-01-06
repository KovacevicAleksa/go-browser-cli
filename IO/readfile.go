package IO

import (
	"fmt"
	"os"
)

func ReadFile(name string) {
	fmt.Printf("Reading file: %s\n", name)

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", name, err)
		return
	}

	fmt.Printf("File name: %s\n", name)
	fmt.Printf("File size: %d bytes\n", len(data))
	fmt.Printf("File content:\n%s\n", string(data))
}
