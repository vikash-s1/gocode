package originator

import (
	"fmt"
	"mementopattern/internal/memento"
	"strings"
)

// TextEditor represents the originator in the memento pattern
type TextEditor struct {
	content   string
	cursorPos int
	filename  string
}

// NewTextEditor creates a new text editor instance
func NewTextEditor(filename string) *TextEditor {
	return &TextEditor{
		content:   "",
		cursorPos: 0,
		filename:  filename,
	}
}

// Write adds text at the current cursor position
func (te *TextEditor) Write(text string) {
	if te.cursorPos < 0 {
		te.cursorPos = 0
	}
	if te.cursorPos > len(te.content) {
		te.cursorPos = len(te.content)
	}

	// Insert text at cursor position
	before := te.content[:te.cursorPos]
	after := te.content[te.cursorPos:]
	te.content = before + text + after
	te.cursorPos += len(text)

	fmt.Printf("✏️  Wrote: '%s' at position %d\n", text, te.cursorPos-len(text))
}

// Delete removes specified number of characters before cursor
func (te *TextEditor) Delete(count int) {
	if count <= 0 || te.cursorPos == 0 {
		return
	}

	startPos := te.cursorPos - count
	if startPos < 0 {
		startPos = 0
	}

	deletedText := te.content[startPos:te.cursorPos]
	te.content = te.content[:startPos] + te.content[te.cursorPos:]
	te.cursorPos = startPos

	fmt.Printf("🗑️  Deleted: '%s' (removed %d characters)\n", deletedText, len(deletedText))
}

// SetCursor sets the cursor position
func (te *TextEditor) SetCursor(pos int) {
	if pos < 0 {
		pos = 0
	}
	if pos > len(te.content) {
		pos = len(te.content)
	}

	te.cursorPos = pos
	fmt.Printf("📍 Cursor moved to position %d\n", pos)
}

// Replace replaces text in a specific range
func (te *TextEditor) Replace(start, end int, newText string) {
	if start < 0 {
		start = 0
	}
	if end > len(te.content) {
		end = len(te.content)
	}
	if start > end {
		start, end = end, start
	}

	oldText := te.content[start:end]
	te.content = te.content[:start] + newText + te.content[end:]
	te.cursorPos = start + len(newText)

	fmt.Printf("🔄 Replaced: '%s' with '%s'\n", oldText, newText)
}

// Clear clears all content
func (te *TextEditor) Clear() {
	te.content = ""
	te.cursorPos = 0
	fmt.Println("🧹 Content cleared")
}

// GetContent returns the current content
func (te *TextEditor) GetContent() string {
	return te.content
}

// GetCursorPos returns the current cursor position
func (te *TextEditor) GetCursorPos() int {
	return te.cursorPos
}

// GetFilename returns the filename
func (te *TextEditor) GetFilename() string {
	return te.filename
}

// GetStats returns editor statistics
func (te *TextEditor) GetStats() (int, int, int) {
	lines := strings.Count(te.content, "\n") + 1
	if te.content == "" {
		lines = 0
	}
	words := len(strings.Fields(te.content))
	chars := len(te.content)
	return lines, words, chars
}

// CreateMemento creates a snapshot of the current state
func (te *TextEditor) CreateMemento(description string) *memento.TextEditorMemento {
	fmt.Printf("💾 Creating memento: %s\n", description)
	return memento.NewTextEditorMemento(te.content, te.cursorPos, description)
}

// RestoreFromMemento restores state from a memento
func (te *TextEditor) RestoreFromMemento(m *memento.TextEditorMemento) {
	te.content = m.GetContent()
	te.cursorPos = m.GetCursorPos()
	fmt.Printf("📂 Restored from memento: %s (created at %s)\n", 
		m.GetDescription(), m.GetTimestamp().Format("15:04:05"))
}

// Display shows the current editor state
func (te *TextEditor) Display() {
	fmt.Println("\n📄 Text Editor State:")
	fmt.Printf("File: %s\n", te.filename)
	fmt.Printf("Cursor Position: %d\n", te.cursorPos)
	
	lines, words, chars := te.GetStats()
	fmt.Printf("Stats: %d lines, %d words, %d characters\n", lines, words, chars)
	
	fmt.Println("Content:")
	if te.content == "" {
		fmt.Println("  [Empty]")
	} else {
		// Show content with cursor indicator
		content := te.content
		if te.cursorPos <= len(content) {
			content = content[:te.cursorPos] + "|" + content[te.cursorPos:]
		}
		
		// Display with line numbers
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			fmt.Printf("  %2d: %s\n", i+1, line)
		}
	}
	fmt.Println()
}