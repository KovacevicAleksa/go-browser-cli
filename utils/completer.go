package utils

import (
	"github.com/chzyer/readline"
)

// Function to create an auto-completer for commands
func Completer(commands []string) *readline.PrefixCompleter {
	var items []readline.PrefixCompleterInterface
	for _, cmd := range commands {
		items = append(items, readline.PcItem(cmd))
	}
	return readline.NewPrefixCompleter(items...)
}
