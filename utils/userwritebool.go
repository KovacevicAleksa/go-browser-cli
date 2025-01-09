package utils

import (
	"fmt"
	"strings"
)

func UserWriteBool(prompt string) bool {
	fmt.Println(prompt)

	var input string
	fmt.Scanln(&input)

	return strings.ToLower(input) == "true"
}
