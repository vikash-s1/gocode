package components

import (
	"fmt"
)

// ModeratorUser represents a user with moderator privileges
type ModeratorUser struct {
	*BaseUser
}

// NewModeratorUser creates a new ModeratorUser
func NewModeratorUser(name string) *ModeratorUser {
	return &ModeratorUser{
		BaseUser: NewBaseUser(name),
	}
}

// SendAnnouncement sends an announcement to all users
func (mu *ModeratorUser) SendAnnouncement(message string) {
	if mu.mediator != nil {
		fmt.Printf("📢 Moderator %s announces: %s\n", mu.name, message)
		mu.mediator.Notify(mu, "broadcast_announcement", message)
	}
}

// HandleNotification handles notifications with moderator-specific behavior
func (mu *ModeratorUser) HandleNotification(event string, data interface{}) {
	switch event {
	case "user_joined":
		username := data.(string)
		fmt.Printf("🛡️  Moderator %s welcomes %s to the chat\n", mu.name, username)
	case "user_left":
		username := data.(string)
		fmt.Printf("🛡️  Moderator %s notes that %s left the chat\n", mu.name, username)
	default:
		// Handle other notifications normally
		mu.BaseUser.HandleNotification(event, data)
	}
}