package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserWriteString(text string) string {
	if text != "" {
		fmt.Println(text)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimRight(input, "\n")

	return input
}
