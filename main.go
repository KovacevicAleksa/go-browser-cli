package main

import (
	"fmt"
	boot "go-browser/boot"
	commands "go-browser/commands"
	utils "go-browser/utils"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

type CommandHandler struct {
	command     string
	handler     func()
	description string
}

func main() {
	// Initialize the bootloader
	boot.BootLoader()

	// Define the commands and their handlers
	commandHandlers := []CommandHandler{
		{"/help", commands.HandleHelp, "Displays the list of available commands"},
		{"/exit", commands.HandleExit, "Exits the application"},
		{"/create", commands.HandleCreate, "Creates a new item"},
		{"/read", commands.HandleRead, "Reads an item"},
		{"/delete", commands.HandleDelete, "Deletes an item"},
		{"/update", commands.HandleUpdate, "Updates an item"},
		{"/rename", commands.HandleRename, "Renames an item"},
		{"/about", commands.HandleAbout, "Displays information about the application"},
		{"/list", commands.HandleList, "Lists all items"},
		{"/aichat", commands.HandleAIChat, "Starts an AI chat session"},
		{"/google", commands.HandleGoogle, "Searches Google"},
		{"/siteperformance", commands.HandleSitePerformance, "Analyzes site performance"},
		{"/sitecontent", commands.HandleSiteContent, "Analyzes site content"},
	}

	// Generate the command list for autocompletion
	commandsList := make([]string, len(commandHandlers))
	commandMap := make(map[string]func())
	for i, cmd := range commandHandlers {
		commandsList[i] = cmd.command
		commandMap[cmd.command] = cmd.handler
	}

	// Initialize readline
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "Enter command: ",
		AutoComplete:    utils.Completer(commandsList),
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
			printHelp(commandHandlers)
		} else {
			fmt.Println("Invalid command. Type /help for a list of available commands.")
		}
	}
}

// printHelp displays the list of available commands and their descriptions
func printHelp(commands []CommandHandler) {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("%s - %s\n", cmd.command, cmd.description)
	}
}
