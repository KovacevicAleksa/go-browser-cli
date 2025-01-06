package main

import (
	IO "go-browser/IO"
)

func main() {

	// name string, text string
	IO.CreateFile(IO.UserWrite("enter file name"), IO.UserWrite("enter file text"))
	// name string (with file type)
	IO.ReadFile(IO.UserWrite("enter file name for read"))
}
