package main

import (
	"fmt"
	IO "go-browser/IO"
)

func main() {
	for {
		command := IO.UserWriteString("Enter command (/help for assistance, /exit to quit):")

		if command == "/exit" {
			fmt.Println("Exiting program.")
			break
		}

		switch command {
		case "/create":
			name := IO.UserWriteString("Enter file name:")
			text := IO.UserWriteString("Enter file content:")
			IO.CreateFile(name, text)
		case "/read":
			name := IO.UserWriteString("Enter file name for reading:")
			IO.ReadFile(name)
		case "/delete":
			name := IO.UserWriteString("Enter file name for deletion:")
			IO.DeleteFile(name)
		case "/help":
			page := IO.UserWriteNum("Enter page number for help:")
			IO.DisplayHelp(page)
		default:
			fmt.Println("Invalid command. Type /help for available commands.")
		}
	}
}
