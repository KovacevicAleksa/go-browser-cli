package IO

import (
	"fmt"
)

func UserWriteString(text string) string {

	fmt.Println(text)

	var first string

	fmt.Scanln(&first)

	return first
}
