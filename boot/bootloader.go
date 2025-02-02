package Boot

import (
	"fmt"
	utils "go-browser/utils"
	"log"
)

// BootLoader initializes the system, configures the logger, and creates necessary folders
func BootLoader() {
	// Print a dashed line for a clean log output
	utils.PrintDashedLine()

	// Define the folder name and ensure its existence
	const folderName = "user_files/logs"
	if err := setupFolder(folderName); err != nil {
		log.Fatalf("ERROR: Boot failed: %v", err)
	}

	// Define the file path and ensure its existence
	const filePath = "user_files/history/history.txt"
	if err := setupFile(filePath); err != nil {
		log.Fatalf("ERROR: Boot failed: %v", err)
	}

	// Initialize logger
	configureLogger()

	fmt.Println("INFO: System booted successfully")
}
