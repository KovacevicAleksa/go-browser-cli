package types

// CommandHandler represents a command with its associated handler function and description.
type CommandHandler struct {
	Command     string // The command string, e.g. "/help"
	Handler     func() // The function that will handle the command
	Description string // A description of what the command does
}
