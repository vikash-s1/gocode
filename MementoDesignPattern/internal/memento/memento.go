package memento

import (
	"time"
)

// Memento interface defines the contract for all mementos
type Memento interface {
	GetTimestamp() time.Time
	GetDescription() string
}

// TextEditorMemento stores the state of a text editor
type TextEditorMemento struct {
	content     string
	cursorPos   int
	timestamp   time.Time
	description string
}

// NewTextEditorMemento creates a new text editor memento
func NewTextEditorMemento(content string, cursorPos int, description string) *TextEditorMemento {
	return &TextEditorMemento{
		content:     content,
		cursorPos:   cursorPos,
		timestamp:   time.Now(),
		description: description,
	}
}

// GetContent returns the saved content (package-private access)
func (m *TextEditorMemento) GetContent() string {
	return m.content
}

// GetCursorPos returns the saved cursor position
func (m *TextEditorMemento) GetCursorPos() int {
	return m.cursorPos
}

// GetTimestamp returns when the memento was created
func (m *TextEditorMemento) GetTimestamp() time.Time {
	return m.timestamp
}

// GetDescription returns a description of the memento
func (m *TextEditorMemento) GetDescription() string {
	return m.description
}

// GameStateMemento stores the state of a game
type GameStateMemento struct {
	level       int
	score       int
	playerName  string
	lives       int
	timestamp   time.Time
	description string
}

// NewGameStateMemento creates a new game state memento
func NewGameStateMemento(level, score int, playerName string, lives int, description string) *GameStateMemento {
	return &GameStateMemento{
		level:       level,
		score:       score,
		playerName:  playerName,
		lives:       lives,
		timestamp:   time.Now(),
		description: description,
	}
}

// GetLevel returns the saved level
func (m *GameStateMemento) GetLevel() int {
	return m.level
}

// GetScore returns the saved score
func (m *GameStateMemento) GetScore() int {
	return m.score
}

// GetPlayerName returns the saved player name
func (m *GameStateMemento) GetPlayerName() string {
	return m.playerName
}

// GetLives returns the saved lives count
func (m *GameStateMemento) GetLives() int {
	return m.lives
}

// GetTimestamp returns when the memento was created
func (m *GameStateMemento) GetTimestamp() time.Time {
	return m.timestamp
}

// GetDescription returns a description of the memento
func (m *GameStateMemento) GetDescription() string {
	return m.description
}