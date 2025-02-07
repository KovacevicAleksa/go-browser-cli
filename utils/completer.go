package utils

import (
	"errors"

	"github.com/chzyer/readline"
)

func Completer(commands []string) (*readline.PrefixCompleter, error) {
	if len(commands) == 0 {
		return nil, errors.New("commands slice cannot be empty")
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
