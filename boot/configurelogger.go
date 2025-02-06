package boot

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// configureLogger sets up the logger with file rotation and log levels
func configureLogger() {
	fileLogger := &lumberjack.Logger{
		Filename:   "user_files/logs/system.log", // Log file name
		MaxSize:    10,                           // Max size in MB before rotation
		MaxBackups: 5,                            // Max number of old log files to retain
		MaxAge:     30,                           // Max age in days to retain old logs
		Compress:   true,                         // Compress rotated files
	}

	// Create a multi-writer to write to both console and file
	multiWriter := io.MultiWriter(os.Stdout, fileLogger)
	log.SetOutput(multiWriter)

	log.Println("INFO: Logger configured successfully")
}
