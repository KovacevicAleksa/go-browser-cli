package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	AI "go-browser/AI"
	IO "go-browser/IO"
	Boot "go-browser/boot"
	searchbrowser "go-browser/search-browser"
	"go-browser/site"
	Site "go-browser/site"
	utils "go-browser/utils"

	"github.com/chzyer/readline"
)

func main() {
	// Initialize the bootloader
	Boot.BootLoader()

	// List of available commands
	commands := []string{
		"/help", "/exit", "/create", "/read", "/delete", "/update",
		"/rename", "/about", "/list", "/aichat", "/google",
		"/siteperformance", "/sitecontent",
	}

	// Initialize readline
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "Enter command: ",
		AutoComplete:    utils.Completer(commands),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("Failed to initialize readline: %v", err)
	}
	defer rl.Close()

	for {
		// Print a dashed line for separation
		utils.PrintDashedLine()

		fmt.Println("Enter command /help for help, use tab for suggestion")

		// Read user input
		command, err := rl.Readline()
		if err != nil {
			fmt.Println("Error reading input. Exiting program.")
			break
		}
		command = strings.TrimSpace(command)

		// Handle different commands
		switch command {
		case "/help":
			page := utils.UserWriteNum("Enter page number for help (e.g., 1, 2, 3):")
			utils.DisplayHelp(page)

		case "/exit":
			fmt.Println("Exiting program.")
			return

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

		case "/update":
			name := utils.UserWriteString("Enter file name for update:")
			text := utils.UserWriteString("Enter new content:")
			IO.UpdateFile(name, text)

		case "/rename":
			name := utils.UserWriteString("Enter file name to rename:")
			newName := utils.UserWriteString("Enter new file name:")
			IO.RenameFile(name, newName)

		case "/about":
			fmt.Println("Version 0.0.0 - Go Browser Tool")

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
			} else {
				fmt.Println(result)
			}

		case "/siteperformance":
			url := utils.UserWriteString("Enter site URL to test performance:")
			timeout := 10 * time.Second
			err := Site.MeasureSitePerformance(url, timeout)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "/sitecontent":
			url := utils.UserWriteString("Enter site URL:")
			element := utils.UserWriteString("Specify the target element (or leave empty for all):")
			includeAttributes := utils.UserWriteBool("(true/false) Include attributes in HTML elements?")
			filter := utils.UserWriteBool("(true/false) Filter unnecessary elements like scripts?")

			content, err := site.SiteContent(url, element, includeAttributes, filter)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Extracted Content:")
				fmt.Println(content)

				save := utils.UserWriteBool("(true/false) Save content?")
				if save {
					name := utils.UserWriteString("Enter file name to save content:")
					IO.CreateFile(name, content)
				}
			}

		default:
			fmt.Println("Invalid command. Type /help for a list of available commands.")
		}
	}
}
