package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserWriteJson prompts the user to input a JSON string. It validates the input to ensure it's a valid JSON format.
// The function will continue to ask for input until a valid JSON string is provided, then it returns the valid JSON string.
func UserWriteJson(text string) string {
	var input string

	for {
		fmt.Println(text)
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Println("ERROR: Failed to read input:", err)
			continue
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal([]byte(input), &jsonData)

		if err != nil {
			log.Println("WARN: Invalid JSON format. Please try again.")
		} else {
			// If input is valid JSON, return it as a string
			return input
		}
	}
}
