package originator

import (
	"fmt"
	"mementopattern/internal/memento"
)

// GameState represents a game's current state (another originator example)
type GameState struct {
	level      int
	score      int
	playerName string
	lives      int
	gameTitle  string
}

// NewGameState creates a new game state instance
func NewGameState(gameTitle, playerName string) *GameState {
	return &GameState{
		level:     1,
		score:     0,
		playerName: playerName,
		lives:     3,
		gameTitle: gameTitle,
	}
}

// LevelUp increases the level and score
func (gs *GameState) LevelUp() {
	gs.level++
	gs.score += gs.level * 100
	fmt.Printf("🎮 Level up! Now at level %d (Score: %d)\n", gs.level, gs.score)
}

// AddScore adds points to the current score
func (gs *GameState) AddScore(points int) {
	gs.score += points
	fmt.Printf("⭐ Added %d points! Total score: %d\n", points, gs.score)
}

// LoseLife decreases lives by one
func (gs *GameState) LoseLife() {
	if gs.lives > 0 {
		gs.lives--
		fmt.Printf("💔 Lost a life! Lives remaining: %d\n", gs.lives)
	}
}

// GainLife increases lives by one
func (gs *GameState) GainLife() {
	gs.lives++
	fmt.Printf("❤️  Gained a life! Lives: %d\n", gs.lives)
}

// SetLevel sets the current level
func (gs *GameState) SetLevel(level int) {
	if level < 1 {
		level = 1
	}
	gs.level = level
	fmt.Printf("🎯 Level set to %d\n", level)
}

// ResetGame resets the game to initial state
func (gs *GameState) ResetGame() {
	gs.level = 1
	gs.score = 0
	gs.lives = 3
	fmt.Println("🔄 Game reset to initial state")
}

// GetLevel returns the current level
func (gs *GameState) GetLevel() int {
	return gs.level
}

// GetScore returns the current score
func (gs *GameState) GetScore() int {
	return gs.score
}

// GetPlayerName returns the player name
func (gs *GameState) GetPlayerName() string {
	return gs.playerName
}

// GetLives returns the current lives
func (gs *GameState) GetLives() int {
	return gs.lives
}

// GetGameTitle returns the game title
func (gs *GameState) GetGameTitle() string {
	return gs.gameTitle
}

// IsGameOver checks if the game is over
func (gs *GameState) IsGameOver() bool {
	return gs.lives <= 0
}

// CreateMemento creates a snapshot of the current game state
func (gs *GameState) CreateMemento(description string) *memento.GameStateMemento {
	fmt.Printf("💾 Creating game save: %s\n", description)
	return memento.NewGameStateMemento(gs.level, gs.score, gs.playerName, gs.lives, description)
}

// RestoreFromMemento restores game state from a memento
func (gs *GameState) RestoreFromMemento(m *memento.GameStateMemento) {
	gs.level = m.GetLevel()
	gs.score = m.GetScore()
	gs.playerName = m.GetPlayerName()
	gs.lives = m.GetLives()
	fmt.Printf("📂 Game loaded: %s (saved at %s)\n", 
		m.GetDescription(), m.GetTimestamp().Format("15:04:05"))
}

// Display shows the current game state
func (gs *GameState) Display() {
	fmt.Println("\n🎮 Game State:")
	fmt.Printf("Game: %s\n", gs.gameTitle)
	fmt.Printf("Player: %s\n", gs.playerName)
	fmt.Printf("Level: %d\n", gs.level)
	fmt.Printf("Score: %d\n", gs.score)
	fmt.Printf("Lives: %d", gs.lives)
	
	// Add visual representation of lives
	fmt.Print(" [")
	for i := 0; i < gs.lives; i++ {
		fmt.Print("❤️ ")
	}
	fmt.Print("]\n")
	
	if gs.IsGameOver() {
		fmt.Println("💀 GAME OVER!")
	}
	fmt.Println()
}