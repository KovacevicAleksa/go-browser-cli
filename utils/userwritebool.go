package utils

import (
	"fmt"
	"strings"
)

func UserWriteBool(prompt string) bool {
	for {
		fmt.Println(prompt)

		var input string
		fmt.Scanln(&input)

		// Normalize input to lowercase
		normalizedInput := strings.ToLower(strings.TrimSpace(input))

		// Check if the input is valid
		if normalizedInput == "true" {
			return true
		} else if normalizedInput == "false" {
			return false
		}

		// If input is invalid, inform the user and repeat the prompt
		fmt.Println("Invalid input. Please enter 'true' or 'false'.")
	}
}
