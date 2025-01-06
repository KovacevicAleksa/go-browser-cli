package utils

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// PrintDashedLine prints a line of dashes that matches the width of the terminal.
func PrintDashedLine() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Greška pri dobijanju širine terminala:", err)

	}

	horizontalLine := ""
	for i := 0; i < width; i++ {
		horizontalLine += "-"
	}

	fmt.Println(horizontalLine)
}
