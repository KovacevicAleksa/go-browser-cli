package utils

import (
	"errors"

	"github.com/chzyer/readline"
)

// Function to create an auto-completer for commands
func Completer(commands []string) (*readline.PrefixCompleter, error) {
	// Error handler: Check if the input slice is nil or empty
	if commands == nil || len(commands) == 0 {
		return nil, errors.New("commands slice cannot be nil or empty")
	}

	var items []readline.PrefixCompleterInterface
	for _, cmd := range commands {
		if cmd == "" {
			return nil, errors.New("commands slice contains an empty string")
		}
		items = append(items, readline.PcItem(cmd))
	}

	return readline.NewPrefixCompleter(items...), nil
}
