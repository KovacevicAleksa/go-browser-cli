package IO

import (
	"fmt"
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
			fmt.Println("Not a valide number")
		}
	}
	return number
}
