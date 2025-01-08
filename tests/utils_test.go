package main

import (
	"bytes"
	"go-browser/utils"
	"os"
	"testing"
)

// TestDisplayHelp tests the DisplayHelp function for proper output based on page input.
func TestDisplayHelp(t *testing.T) {
	page := 1
	expectedOut := "Available commands (Page 1):\n" +
		"/create - Create a new file\n" +
		"/read - Read a file\n" +
		"/delete - Delete a file\n" +
		"/help - Show available commands\n" +
		"/about - Show version\n\n" +
		"Type 'next' to see more commands.\n"

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Call the function
	utils.DisplayHelp(page)

	// Capture the output
	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Error reading from pipe: %v", err)
	}

	// Check if the output matches the expected value
	if buf.String() != expectedOut {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOut, buf.String())
	}
}

// TestUserWriteNum tests UserWriteNum function by simulating user input.
func TestUserWriteNum(t *testing.T) {
	input := "42\n"
	expected := 42

	// Simulate user input
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }() // Ensure to restore stdin after the test
	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	// Call the function and check the result
	actual := utils.UserWriteNum("Enter a number:")
	if actual != expected {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}

// TestUserWriteString tests UserWriteString function by simulating user input.
func TestUserWriteString(t *testing.T) {
	input := "hello\n"
	expected := "hello"

	// Simulate user input
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }() // Ensure to restore stdin after the test
	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	// Call the function and check the result
	actual := utils.UserWriteString("Enter a string:")
	if actual != expected {
		t.Errorf("Expected: %s, Got: %s", expected, actual)
	}
}
