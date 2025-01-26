package utils

import (
	"os"
	"path/filepath"
)

// FolderSize calculates the total size of files within the specified folder.
func FolderSize(folderPath string) (int64, error) {
	var totalSize int64

	// Walk through the directory and accumulate file sizes.
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If it's a file, add its size to the total.
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return totalSize, nil
}
