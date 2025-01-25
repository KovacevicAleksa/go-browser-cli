package Boot

import (
	"log"

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
	log.Println("INFO: Logger configured successfully")

}
