package main

import (
	"fmt"
	AI "go-browser/AI"
	IO "go-browser/IO"
	Boot "go-browser/boot"
	utils "go-browser/utils"
)

func main() {

	Boot.BootLoader()

	for {
		utils.PrintDashedLine()

		command := utils.UserWriteString("Enter command (/help for assistance, /exit to quit):")

		if command == "/exit" {
			fmt.Println("Exiting program.")
			break
		}

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
			about := "Ver v0.0.0"
			fmt.Println(about)
		case "/list":
			IO.ListFile(".")
		case "/aichat":
			text := utils.UserWriteString("Enter text")
			response := AI.ChatGPT(text)
			fmt.Println(response)
		default:
			fmt.Println("Invalid command. Type /help for available commands.")
		}
	}
}
