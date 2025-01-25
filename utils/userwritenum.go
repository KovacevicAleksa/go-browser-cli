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
		fmt.Scanln(&input)

		// conver string to num
		num, err := strconv.Atoi(input)
		if err == nil {
			number = num
			break
		} else {
			log.Println("WARN: Not a valide number")
		}
	}
	return number
}
