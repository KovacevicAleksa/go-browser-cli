package IO

import (
	"fmt"
	"os"
)

func CreateFile(name string, text string) {
	fmt.Println("Creating file")
	file, err := os.Create(name)

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString(text)

	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s", file.Name())
	fmt.Printf("\nfile length: %d\n", length)

}
