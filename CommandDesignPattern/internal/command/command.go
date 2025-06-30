package command

// Command interface defines the contract for all commands
type Command interface {
	Execute() error
	Undo() error
	GetDescription() string
}

// UndoableCommand extends Command with undo capability tracking
type UndoableCommand interface {
	Command
	CanUndo() bool
}

// NoOpCommand represents a null command (Null Object pattern)
type NoOpCommand struct{}

func NewNoOpCommand() *NoOpCommand {
	return &NoOpCommand{}
}

func (n *NoOpCommand) Execute() error {
	return nil
}

func (n *NoOpCommand) Undo() error {
	return nil
}

func (n *NoOpCommand) GetDescription() string {
	return "No Operation"
}

func (n *NoOpCommand) CanUndo() bool {
	return false
}