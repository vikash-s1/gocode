package commands

import (
	"commandpattern/internal/command"
	"fmt"
	"strings"
)

// MacroCommand executes multiple commands as one
type MacroCommand struct {
	commands []command.Command
	name     string
}

func NewMacroCommand(name string, commands []command.Command) command.Command {
	return &MacroCommand{
		commands: commands,
		name:     name,
	}
}

func (m *MacroCommand) Execute() error {
	fmt.Printf("🎯 Executing macro: %s\n", m.name)
	
	for i, cmd := range m.commands {
		fmt.Printf("  Step %d: %s\n", i+1, cmd.GetDescription())
		if err := cmd.Execute(); err != nil {
			return fmt.Errorf("macro command failed at step %d: %w", i+1, err)
		}
	}
	
	fmt.Printf("✅ Macro '%s' completed successfully\n", m.name)
	return nil
}

func (m *MacroCommand) Undo() error {
	fmt.Printf("↩️  Undoing macro: %s\n", m.name)
	
	// Undo commands in reverse order
	for i := len(m.commands) - 1; i >= 0; i-- {
		cmd := m.commands[i]
		fmt.Printf("  Undoing step %d: %s\n", i+1, cmd.GetDescription())
		if err := cmd.Undo(); err != nil {
			return fmt.Errorf("macro undo failed at step %d: %w", i+1, err)
		}
	}
	
	fmt.Printf("✅ Macro '%s' undo completed\n", m.name)
	return nil
}

func (m *MacroCommand) GetDescription() string {
	var descriptions []string
	for _, cmd := range m.commands {
		descriptions = append(descriptions, cmd.GetDescription())
	}
	return fmt.Sprintf("Macro '%s': [%s]", m.name, strings.Join(descriptions, ", "))
}

func (m *MacroCommand) CanUndo() bool {
	// Macro can be undone if all its commands can be undone
	for _, cmd := range m.commands {
		if undoable, ok := cmd.(command.UndoableCommand); ok {
			if !undoable.CanUndo() {
				return false
			}
		}
	}
	return true
}