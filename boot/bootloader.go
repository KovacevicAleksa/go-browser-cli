package Boot

import (
	"fmt"
	utils "go-browser/utils"
	"os"
)

func BootLoader() {

	utils.PrintDashedLine()

	println("Booting system...")
	// Define the folder name
	folderName := "user_files"

	// Ensure the folder exists
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating folder %s: %v\n", folderName, err)
		return
	}

	fmt.Printf("Main Folder created successfully: [%s]\n", folderName)
}
