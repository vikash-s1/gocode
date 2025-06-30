package invoker

import (
	"commandpattern/internal/command"
	"fmt"
	"strings"
)

// RemoteControl acts as the invoker in the command pattern
type RemoteControl struct {
	onCommands  []command.Command
	offCommands []command.Command
	undoCommand command.Command
	history     []command.Command
}

func NewRemoteControl() *RemoteControl {
	noOpCommand := command.NewNoOpCommand()
	
	// Initialize with 7 slots
	onCommands := make([]command.Command, 7)
	offCommands := make([]command.Command, 7)
	
	// Fill with NoOp commands initially
	for i := 0; i < 7; i++ {
		onCommands[i] = noOpCommand
		offCommands[i] = noOpCommand
	}
	
	return &RemoteControl{
		onCommands:  onCommands,
		offCommands: offCommands,
		undoCommand: noOpCommand,
		history:     make([]command.Command, 0),
	}
}

// SetCommand sets both on and off commands for a slot
func (r *RemoteControl) SetCommand(slot int, onCommand, offCommand command.Command) {
	if slot < 0 || slot >= len(r.onCommands) {
		fmt.Printf("❌ Invalid slot number: %d\n", slot)
		return
	}
	
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
	fmt.Printf("✅ Slot %d configured: ON=%s, OFF=%s\n", 
		slot, onCommand.GetDescription(), offCommand.GetDescription())
}

// OnButtonPressed executes the on command for the given slot
func (r *RemoteControl) OnButtonPressed(slot int) {
	if slot < 0 || slot >= len(r.onCommands) {
		fmt.Printf("❌ Invalid slot number: %d\n", slot)
		return
	}
	
	cmd := r.onCommands[slot]
	fmt.Printf("🔘 ON button pressed for slot %d\n", slot)
	
	if err := cmd.Execute(); err != nil {
		fmt.Printf("❌ Command execution failed: %v\n", err)
		return
	}
	
	r.undoCommand = cmd
	r.addToHistory(cmd)
}

// OffButtonPressed executes the off command for the given slot
func (r *RemoteControl) OffButtonPressed(slot int) {
	if slot < 0 || slot >= len(r.offCommands) {
		fmt.Printf("❌ Invalid slot number: %d\n", slot)
		return
	}
	
	cmd := r.offCommands[slot]
	fmt.Printf("🔘 OFF button pressed for slot %d\n", slot)
	
	if err := cmd.Execute(); err != nil {
		fmt.Printf("❌ Command execution failed: %v\n", err)
		return
	}
	
	r.undoCommand = cmd
	r.addToHistory(cmd)
}

// UndoButtonPressed undoes the last command
func (r *RemoteControl) UndoButtonPressed() {
	fmt.Println("↩️  UNDO button pressed")
	
	if err := r.undoCommand.Undo(); err != nil {
		fmt.Printf("❌ Undo failed: %v\n", err)
		return
	}
	
	fmt.Println("✅ Last command undone successfully")
}

// ExecuteCommand directly executes any command
func (r *RemoteControl) ExecuteCommand(cmd command.Command) {
	fmt.Printf("🎯 Executing command: %s\n", cmd.GetDescription())
	
	if err := cmd.Execute(); err != nil {
		fmt.Printf("❌ Command execution failed: %v\n", err)
		return
	}
	
	r.undoCommand = cmd
	r.addToHistory(cmd)
	fmt.Println("✅ Command executed successfully")
}

// addToHistory adds a command to the history
func (r *RemoteControl) addToHistory(cmd command.Command) {
	r.history = append(r.history, cmd)
	
	// Keep only last 10 commands in history
	if len(r.history) > 10 {
		r.history = r.history[1:]
	}
}

// ShowHistory displays the command history
func (r *RemoteControl) ShowHistory() {
	fmt.Println("\n📜 Command History:")
	fmt.Println(strings.Repeat("-", 40))
	
	if len(r.history) == 0 {
		fmt.Println("No commands executed yet")
		return
	}
	
	for i, cmd := range r.history {
		fmt.Printf("%d. %s\n", i+1, cmd.GetDescription())
	}
}

// ShowConfiguration displays the current remote configuration
func (r *RemoteControl) ShowConfiguration() {
	fmt.Println("\n🎛️  Remote Control Configuration:")
	fmt.Println(strings.Repeat("=", 50))
	
	for i := 0; i < len(r.onCommands); i++ {
		onDesc := r.onCommands[i].GetDescription()
		offDesc := r.offCommands[i].GetDescription()
		
		if onDesc != "No Operation" || offDesc != "No Operation" {
			fmt.Printf("Slot %d: ON=[%s] | OFF=[%s]\n", i, onDesc, offDesc)
		} else {
			fmt.Printf("Slot %d: [Empty]\n", i)
		}
	}
	
	fmt.Printf("\nLast Command: %s\n", r.undoCommand.GetDescription())
}

