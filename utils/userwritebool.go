package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// UserWriteBool prompts the user for a 'true' or 'false' input and returns the corresponding boolean value.
// It will continue to ask for input until a valid response is provided.
func UserWriteBool(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("ERROR: Failed to read input:", err)
			continue
		}

		// Normalize input to lowercase
		normalizedInput := strings.ToLower(strings.TrimSpace(input))

		// Check if the input is valid
		if normalizedInput == "true" {
			return true
		} else if normalizedInput == "false" {
			return false
		}

		// If input is invalid, inform the user and repeat the prompt
		log.Println("WARN: Invalid input. Please enter 'true' or 'false'.")
	}
}
