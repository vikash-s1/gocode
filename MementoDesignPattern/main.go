package main

import (
	"fmt"
	"mementopattern/internal/caretaker"
	"mementopattern/internal/memento"
	"mementopattern/internal/originator"
	"strings"
)

func main() {
	fmt.Println("=== Memento Design Pattern Demo ===")
	fmt.Println("Text Editor with Undo/Redo Functionality")
	fmt.Println(strings.Repeat("=", 60))

	// Demo 1: Text Editor with Memento Pattern
	textEditorDemo()
	
	fmt.Println(strings.Repeat("=", 60))
	
	// Demo 2: Game State with Memento Pattern
	gameStateDemo()
}

func textEditorDemo() {
	fmt.Println("\n1. Text Editor Demo")
	fmt.Println(strings.Repeat("-", 40))

	// Create text editor and history manager
	editor := originator.NewTextEditor("document.txt")
	history := caretaker.NewHistoryManager("Text Editor", 5)

	// Initial state
	editor.Display()
	history.SaveMemento(editor.CreateMemento("Initial empty document"))

	// Step 1: Write some text
	fmt.Println("Step 1: Writing initial text")
	editor.Write("Hello, World!")
	editor.Display()
	history.SaveMemento(editor.CreateMemento("Added greeting"))

	// Step 2: Add more text
	fmt.Println("Step 2: Adding more content")
	editor.Write("\nThis is a demo of the Memento pattern.")
	editor.Display()
	history.SaveMemento(editor.CreateMemento("Added description"))

	// Step 3: Move cursor and insert text
	fmt.Println("Step 3: Inserting text in middle")
	editor.SetCursor(7) // After "Hello, "
	editor.Write("beautiful ")
	editor.Display()
	history.SaveMemento(editor.CreateMemento("Inserted adjective"))

	// Step 4: Delete some text
	fmt.Println("Step 4: Deleting text")
	editor.SetCursor(len(editor.GetContent()))
	editor.Delete(8) // Delete " pattern"
	editor.Display()
	history.SaveMemento(editor.CreateMemento("Removed word"))

	// Show history
	history.ShowHistory()

	// Step 5: Undo operations
	fmt.Println("Step 5: Testing Undo functionality")
	if m, ok := history.Undo(); ok {
		if textMemento, ok := m.(*memento.TextEditorMemento); ok {
			editor.RestoreFromMemento(textMemento)
			editor.Display()
		}
	}

	if m, ok := history.Undo(); ok {
		if textMemento, ok := m.(*memento.TextEditorMemento); ok {
			editor.RestoreFromMemento(textMemento)
			editor.Display()
		}
	}

	// Step 6: Redo operations
	fmt.Println("Step 6: Testing Redo functionality")
	if m, ok := history.Redo(); ok {
		if textMemento, ok := m.(*memento.TextEditorMemento); ok {
			editor.RestoreFromMemento(textMemento)
			editor.Display()
		}
	}

	// Step 7: Jump to specific point in history
	fmt.Println("Step 7: Jumping to specific point in history")
	if m, ok := history.JumpToMemento(1); ok { // Jump to second state
		if textMemento, ok := m.(*memento.TextEditorMemento); ok {
			editor.RestoreFromMemento(textMemento)
			editor.Display()
		}
	}

	history.ShowHistory()
}

func gameStateDemo() {
	fmt.Println("\n2. Game State Demo")
	fmt.Println(strings.Repeat("-", 40))

	// Create game and history manager
	game := originator.NewGameState("Super Adventure", "Player1")
	gameHistory := caretaker.NewHistoryManager("Game Save", 10)

	// Initial state
	game.Display()
	gameHistory.SaveMemento(game.CreateMemento("Game Start"))

	// Step 1: Play and level up
	fmt.Println("Step 1: Playing the game")
	game.AddScore(150)
	game.LevelUp()
	game.Display()
	gameHistory.SaveMemento(game.CreateMemento("Level 2 Checkpoint"))

	// Step 2: Continue playing
	fmt.Println("Step 2: More progress")
	game.AddScore(200)
	game.LevelUp()
	game.GainLife()
	game.Display()
	gameHistory.SaveMemento(game.CreateMemento("Level 3 with Extra Life"))

	// Step 3: Difficult section - lose lives
	fmt.Println("Step 3: Difficult section")
	game.AddScore(50)
	game.LoseLife()
	game.LoseLife()
	game.Display()
	gameHistory.SaveMemento(game.CreateMemento("After Difficult Section"))

	// Step 4: Boss fight - things go wrong
	fmt.Println("Step 4: Boss fight gone wrong")
	game.LoseLife()
	game.Display()

	if game.IsGameOver() {
		fmt.Println("💀 Game Over! Let's restore from a checkpoint...")
		
		// Restore from previous checkpoint
		if m, ok := gameHistory.Undo(); ok {
			if gameMemento, ok := m.(*memento.GameStateMemento); ok {
				game.RestoreFromMemento(gameMemento)
				game.Display()
			}
		}
	}

	// Step 5: Try again with better strategy
	fmt.Println("Step 5: Second attempt at boss fight")
	game.AddScore(500) // Successful boss fight
	game.LevelUp()
	game.Display()
	gameHistory.SaveMemento(game.CreateMemento("Boss Defeated!"))

	// Show game history
	gameHistory.ShowHistory()

	// Step 6: Load different save points
	fmt.Println("Step 6: Loading different save points")
	
	// Jump to level 2 checkpoint
	if m, ok := gameHistory.JumpToMemento(1); ok {
		if gameMemento, ok := m.(*memento.GameStateMemento); ok {
			game.RestoreFromMemento(gameMemento)
			game.Display()
		}
	}

	// Jump back to final state
	if m, ok := gameHistory.JumpToMemento(gameHistory.GetHistorySize()-1); ok {
		if gameMemento, ok := m.(*memento.GameStateMemento); ok {
			game.RestoreFromMemento(gameMemento)
			game.Display()
		}
	}

	gameHistory.ShowHistory()

	fmt.Println("\n=== Demo Complete ===")
	fmt.Println("The Memento pattern provides:")
	fmt.Println("✅ State capture and restoration")
	fmt.Println("✅ Undo/Redo functionality")
	fmt.Println("✅ History management")
	fmt.Println("✅ Encapsulation preservation")
	fmt.Println("✅ Checkpoint/Save game functionality")
	fmt.Println("✅ Rollback capabilities")
}