package utils

import (
	"fmt"
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
		fmt.Println("\nType 'next' to see more commands.")
	} else {
		fmt.Println("\nEnd of list.")
	}
}
