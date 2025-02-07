package utils

import (
	"fmt"
	"log"
	"strconv"
)

// UserWriteNum prompts the user for an integer input and returns the corresponding integer value.
// If the input is not a valid number, it continues to prompt the user until a valid number is provided.
func UserWriteNum(text string) int {
	var input string
	var number int
	for {
		fmt.Println(text)
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Println("ERROR: Failed to read input:", err)
			continue
		}

		// Convert string to number
		num, err := strconv.Atoi(input)
		if err == nil {
			number = num
			break
		} else {
			log.Println("WARN: Not a valid number")
		}
	}
	return number
}
