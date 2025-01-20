package main

import (
	"fmt"
	boot "go-browser/boot"
	"go-browser/commands"
	"go-browser/types"
	"go-browser/utils"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	// Initialize the bootloader
	boot.BootLoader()

	// Define the commands and their handlers
	commandHandlers := []types.CommandHandler{
		{Command: "/help", Handler: commands.HandleHelp, Description: "Displays the list of available commands"},
		{Command: "/exit", Handler: commands.HandleExit, Description: "Exits the application"},
		{Command: "/create", Handler: commands.HandleCreate, Description: "Creates a new file"},
		{Command: "/read", Handler: commands.HandleRead, Description: "Reads an item"},
		{Command: "/delete", Handler: commands.HandleDelete, Description: "Deletes an item"},
		{Command: "/update", Handler: commands.HandleUpdate, Description: "Updates an item"},
		{Command: "/rename", Handler: commands.HandleRename, Description: "Renames an item"},
		{Command: "/about", Handler: commands.HandleAbout, Description: "Displays information about the application"},
		{Command: "/list", Handler: commands.HandleList, Description: "Lists all items"},
		{Command: "/aichat", Handler: commands.HandleAIChat, Description: "Starts an AI chat session"},
		{Command: "/google", Handler: commands.HandleGoogle, Description: "Searches Google"},
		{Command: "/siteperformance", Handler: commands.HandleSitePerformance, Description: "Analyzes site performance"},
		{Command: "/sitecontent", Handler: commands.HandleSiteContent, Description: "Analyzes site content"},
		{Command: "/fetchSiteData", Handler: commands.HandleHttpRequest, Description: "Fetches and analyzes HTTP response data from a given URL"},
	}

	// Generate the command list for autocompletion
	commandsList := make([]string, len(commandHandlers))
	commandMap := make(map[string]func())
	for i, cmd := range commandHandlers {
		commandsList[i] = cmd.Command
		commandMap[cmd.Command] = cmd.Handler
	}

	// Validation check for  Completer func
	autoCompleter, err := utils.Completer(commandsList)
	if err != nil {
		log.Fatalf("Failed to create completer: %v", err)
	}

	// Initialize readline
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "Enter command: ",
		AutoComplete:    autoCompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("Failed to initialize readline: %v", err)
	}
	defer rl.Close()

	// Main loop
	for {
		utils.PrintDashedLine()
		fmt.Println("Enter command /help for help, use tab for suggestion")

		// Read user input
		input, err := rl.Readline()
		if err != nil {
			fmt.Println("Error reading input. Exiting program.")
			break
		}
		command := strings.TrimSpace(input)

		// Handle commands
		if handler, exists := commandMap[command]; exists {
			handler()
		} else if command == "/help" {
			utils.PrintHelp(commandHandlers)
		} else {
			fmt.Println("Invalid command. Type /help for a list of available commands.")
		}
	}
}
