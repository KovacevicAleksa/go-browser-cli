package Boot

import (
	"fmt"
	"os"
)

func setupFolder(folderName string) error {
	_, err := os.Stat(folderName)
	if err == nil {
		// Folder exists, no need to create it
		return nil
	} else if os.IsNotExist(err) {
		// Folder doesn't exist, create it
		err = os.MkdirAll(folderName, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("INFO: User folder %s created successfully\n", folderName)
		return nil
	} else {
		return err
	}
}
