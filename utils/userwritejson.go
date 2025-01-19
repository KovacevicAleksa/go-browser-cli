package utils

import (
	"encoding/json"
	"fmt"
)

func UserWriteJson(text string) string {
	var input string

	for {
		fmt.Println(text)
		fmt.Scanln(&input)

		var jsonData map[string]interface{}
		err := json.Unmarshal([]byte(input), &jsonData)

		if err != nil {
			fmt.Println("Invalid JSON format. Please try again.")
		} else {
			// If input is valid JSON, return it as a string
			return input
		}
	}
}
