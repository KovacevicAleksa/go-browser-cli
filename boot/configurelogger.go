package Boot

import (
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// configureLogger sets up the logger with file rotation and log levels
func configureLogger() {
	// Configure log rotation using Lumberjack
	log.SetOutput(&lumberjack.Logger{
		Filename:   "user_files/logs/system.log", // Log file name
		MaxSize:    10,                           // Max size in MB before rotation
		MaxBackups: 5,                            // Max number of old log files to retain
		MaxAge:     30,                           // Max age in days to retain old logs
		Compress:   true,                         // Compress rotated files
	})

	// Define log levels (INFO, WARN, ERROR)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLog := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Example usage of different log levels
	infoLog.Println("Logger initialized for INFO level")
	warnLog.Println("Logger initialized for WARN level")
	errorLog.Println("Logger initialized for ERROR level")
}
