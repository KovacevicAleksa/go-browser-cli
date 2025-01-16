package utils

import (
	"fmt"
	"go-browser/types"
	"strconv"
)

func DisplayHelp(page int) {
	// Define the list of available commands
	commands := []string{
		"/create - Create a new file",
		"/read - Read a file",
		"/delete - Delete a file",
		"/help - Show available commands",
		"/about - Show version",
		"/list - List all files",
		"/exit - Exit the program",
		"/update - Update a file",
		"/rename - Rename a file",
		"/aichat - AI Chat",
		"/google - Google search",
		"/siteperformance - Test loading time and status code",
		"/sitecontent fetches and cleans HTML content from site",
	}

	// Calculate the start and end indices based on the page
	start := (page - 1) * 5
	end := start + 5

	// Check if the end index is out of bounds
	if end > len(commands) {
		end = len(commands)
	}

	// Display the commands for the current page
	fmt.Println("Available commands (Page " + strconv.Itoa(page) + "):")
	for i := start; i < end; i++ {
		fmt.Println(commands[i])
	}

	// Check if there are more commands to display
	if end < len(commands) {
		fmt.Println("\nFor next page type '/help' and page", page+1)
	} else {
		fmt.Println("\nEnd of list.")
	}
}

// PrintHelp displays the list of available commands and their descriptions.
func PrintHelp(commands []types.CommandHandler) {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		// Print the command and its description
		fmt.Printf("%s - %s\n", cmd.Command, cmd.Description)
	}
}
