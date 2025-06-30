package caretaker

import (
	"fmt"
	"mementopattern/internal/memento"
	"strings"
)

// HistoryManager acts as the caretaker in the memento pattern
type HistoryManager struct {
	mementos    []memento.Memento
	currentPos  int
	maxHistory  int
	description string
}

// NewHistoryManager creates a new history manager
func NewHistoryManager(description string, maxHistory int) *HistoryManager {
	if maxHistory <= 0 {
		maxHistory = 10 // default
	}
	
	return &HistoryManager{
		mementos:    make([]memento.Memento, 0),
		currentPos:  -1,
		maxHistory:  maxHistory,
		description: description,
	}
}

// SaveMemento saves a memento to history
func (hm *HistoryManager) SaveMemento(m memento.Memento) {
	// If we're not at the end of history, remove everything after current position
	if hm.currentPos < len(hm.mementos)-1 {
		hm.mementos = hm.mementos[:hm.currentPos+1]
	}
	
	// Add new memento
	hm.mementos = append(hm.mementos, m)
	hm.currentPos++
	
	// Maintain max history size
	if len(hm.mementos) > hm.maxHistory {
		hm.mementos = hm.mementos[1:]
		hm.currentPos--
	}
	
	fmt.Printf("💾 Saved to history: %s (Position: %d/%d)\n", 
		m.GetDescription(), hm.currentPos+1, len(hm.mementos))
}

// Undo returns the previous memento if available
func (hm *HistoryManager) Undo() (memento.Memento, bool) {
	if !hm.CanUndo() {
		fmt.Println("❌ Cannot undo: No previous state available")
		return nil, false
	}
	
	hm.currentPos--
	m := hm.mementos[hm.currentPos]
	fmt.Printf("↩️  Undo: Restored to '%s' (Position: %d/%d)\n", 
		m.GetDescription(), hm.currentPos+1, len(hm.mementos))
	return m, true
}

// Redo returns the next memento if available
func (hm *HistoryManager) Redo() (memento.Memento, bool) {
	if !hm.CanRedo() {
		fmt.Println("❌ Cannot redo: No next state available")
		return nil, false
	}
	
	hm.currentPos++
	m := hm.mementos[hm.currentPos]
	fmt.Printf("↪️  Redo: Restored to '%s' (Position: %d/%d)\n", 
		m.GetDescription(), hm.currentPos+1, len(hm.mementos))
	return m, true
}

// CanUndo checks if undo is possible
func (hm *HistoryManager) CanUndo() bool {
	return hm.currentPos > 0
}

// CanRedo checks if redo is possible
func (hm *HistoryManager) CanRedo() bool {
	return hm.currentPos < len(hm.mementos)-1
}

// GetCurrentMemento returns the current memento
func (hm *HistoryManager) GetCurrentMemento() (memento.Memento, bool) {
	if hm.currentPos >= 0 && hm.currentPos < len(hm.mementos) {
		return hm.mementos[hm.currentPos], true
	}
	return nil, false
}

// GetMementoAt returns memento at specific position
func (hm *HistoryManager) GetMementoAt(index int) (memento.Memento, bool) {
	if index >= 0 && index < len(hm.mementos) {
		return hm.mementos[index], true
	}
	return nil, false
}

// JumpToMemento jumps to a specific memento in history
func (hm *HistoryManager) JumpToMemento(index int) (memento.Memento, bool) {
	if index >= 0 && index < len(hm.mementos) {
		hm.currentPos = index
		m := hm.mementos[index]
		fmt.Printf("🎯 Jumped to: '%s' (Position: %d/%d)\n", 
			m.GetDescription(), index+1, len(hm.mementos))
		return m, true
	}
	fmt.Printf("❌ Invalid position: %d\n", index)
	return nil, false
}

// Clear clears all history
func (hm *HistoryManager) Clear() {
	hm.mementos = make([]memento.Memento, 0)
	hm.currentPos = -1
	fmt.Println("🧹 History cleared")
}

// GetHistorySize returns the number of mementos in history
func (hm *HistoryManager) GetHistorySize() int {
	return len(hm.mementos)
}

// GetCurrentPosition returns the current position in history
func (hm *HistoryManager) GetCurrentPosition() int {
	return hm.currentPos
}

// ShowHistory displays the complete history
func (hm *HistoryManager) ShowHistory() {
	fmt.Printf("\n📜 %s History:\n", hm.description)
	fmt.Println(strings.Repeat("-", 50))
	
	if len(hm.mementos) == 0 {
		fmt.Println("No history available")
		return
	}
	
	for i, m := range hm.mementos {
		marker := "  "
		if i == hm.currentPos {
			marker = "▶️"
		}
		
		fmt.Printf("%s %d. %s (saved at %s)\n", 
			marker, i+1, m.GetDescription(), m.GetTimestamp().Format("15:04:05"))
	}
	
	fmt.Printf("\nCurrent Position: %d/%d\n", hm.currentPos+1, len(hm.mementos))
	fmt.Printf("Can Undo: %t, Can Redo: %t\n", hm.CanUndo(), hm.CanRedo())
	fmt.Println()
}