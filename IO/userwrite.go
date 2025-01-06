package IO

import (
	"fmt"
)

func UserWrite(text string) string {

	fmt.Println(text)

	var first string

	fmt.Scanln(&first)

	return first
}
