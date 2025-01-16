package utils

import (
	"fmt"
)

func UserWriteString(text string) string {

	if text != "" {
		fmt.Println(text)
	}

	var first string

	fmt.Scanln(&first)

	return first
}
