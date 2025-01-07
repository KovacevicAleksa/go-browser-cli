package main

import (
	"fmt"
	AI "go-browser/AI"
	IO "go-browser/IO"
	Boot "go-browser/boot"
	searchbrowser "go-browser/search-browser"
	utils "go-browser/utils"
)

func main() {
	// Initialize the bootloader
	Boot.BootLoader()

	for {
		utils.PrintDashedLine()

		// Prompt the user for a command
		command := utils.UserWriteString("Enter command (/help for assistance, /exit to quit):")

		// Exit condition
		if command == "/exit" {
			fmt.Println("Exiting program.")
			break
		}

		// Handle different commands
		switch command {
		case "/create":
			name := utils.UserWriteString("Enter file name:")
			text := utils.UserWriteString("Enter file content:")
			IO.CreateFile(name, text)

		case "/read":
			name := utils.UserWriteString("Enter file name for reading:")
			IO.ReadFile(name)

		case "/delete":
			name := utils.UserWriteString("Enter file name for deletion:")
			IO.DeleteFile(name)

		case "/help":
			page := utils.UserWriteNum("Enter page number for help:")
			utils.DisplayHelp(page)

		case "/update":
			name := utils.UserWriteString("Enter file name for update:")
			text := utils.UserWriteString("Enter file content:")
			IO.UpdateFile(name, text)

		case "/rename":
			name := utils.UserWriteString("Enter file name for rename:")
			newname := utils.UserWriteString("Enter new file name for update:")
			IO.RenameFile(name, newname)

		case "/about":
			about := "Version 0.0.0"
			fmt.Println(about)

		case "/list":
			IO.ListFile(".")

		case "/aichat":
			text := utils.UserWriteString("Enter text:")
			response := AI.ChatGPT(text)
			fmt.Println(response)

		case "/google":
			search := utils.UserWriteString("Enter text for search:")
			result, err := searchbrowser.SearchGoogle(search)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(result)

		default:
			fmt.Println("Invalid command. Type /help for available commands.")
		}
	}
}
