package utils

import (
	"fmt"
	"log"
	"strconv"
)

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
