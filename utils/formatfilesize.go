package utils

import (
	"fmt"
)

// FormatFileSize converts a file size in bytes to a human-readable string.
func FormatFileSize(sizeBytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case sizeBytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(sizeBytes)/float64(TB))
	case sizeBytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(sizeBytes)/float64(GB))
	case sizeBytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(sizeBytes)/float64(MB))
	case sizeBytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(sizeBytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", sizeBytes)
	}
}
