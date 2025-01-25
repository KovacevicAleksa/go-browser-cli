package utils

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

// PrintDashedLine prints a line of dashes that matches the width of the terminal.
func PrintDashedLine() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Println("ERROR: Error getting terminal width:", err)

	}

	horizontalLine := ""
	for i := 0; i < width; i++ {
		horizontalLine += "-"
	}

	fmt.Println(horizontalLine)
}
