package Boot

import (
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
		log.Fatalf("Boot failed: %v", err)
	}

	// Initialize logger
	configureLogger()

	log.Printf("INFO: System booted successfully. Folder created: %s", folderName)

}
