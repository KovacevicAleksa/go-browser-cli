package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// UserWriteString prints a message if it's not empty and returns the user input as a string.
func UserWriteString(text string) string {
	if text != "" {
		fmt.Println(text)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimRight(input, "\n")

	return input
}
