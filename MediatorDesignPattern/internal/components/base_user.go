package components

import (
	"fmt"
	"mediator/internal/mediator"
)

// BaseUser provides common functionality for all user types
type BaseUser struct {
	name     string
	mediator mediator.Mediator
	status   string
}

// NewBaseUser creates a new BaseUser
func NewBaseUser(name string) *BaseUser {
	return &BaseUser{
		name:   name,
		status: "online",
	}
}

// SetMediator sets the mediator for this user
func (bu *BaseUser) SetMediator(m mediator.Mediator) {
	bu.mediator = m
}

// GetName returns the user's name
func (bu *BaseUser) GetName() string {
	return bu.name
}

// GetStatus returns the user's current status
func (bu *BaseUser) GetStatus() string {
	return bu.status
}

// SetStatus changes the user's status
func (bu *BaseUser) SetStatus(status string) {
	bu.status = status
	if bu.mediator != nil {
		bu.mediator.Notify(bu, "change_status", status)
	}
}

// SendMessage sends a public message through the mediator
func (bu *BaseUser) SendMessage(message string) {
	if bu.mediator != nil {
		fmt.Printf("💬 %s says: %s\n", bu.name, message)
		bu.mediator.Notify(bu, "send_message", message)
	}
}

// SendPrivateMessage sends a private message to a specific user
func (bu *BaseUser) SendPrivateMessage(recipient, message string) {
	if bu.mediator != nil {
		fmt.Printf("🔒 %s whispers to %s: %s\n", bu.name, recipient, message)
		data := map[string]string{
			"recipient": recipient,
			"message":   message,
		}
		bu.mediator.Notify(bu, "send_private_message", data)
	}
}

// RequestUserList requests the list of active users
func (bu *BaseUser) RequestUserList() {
	if bu.mediator != nil {
		bu.mediator.Notify(bu, "request_user_list", nil)
	}
}

// HandleNotification handles notifications from the mediator
func (bu *BaseUser) HandleNotification(event string, data interface{}) {
	switch event {
	case "message_received":
		msgData := data.(map[string]string)
		fmt.Printf("📨 %s received message from %s: %s\n", bu.name, msgData["sender"], msgData["message"])
	case "private_message_received":
		msgData := data.(map[string]string)
		fmt.Printf("🔐 %s received private message from %s: %s\n", bu.name, msgData["sender"], msgData["message"])
	case "announcement_received":
		msgData := data.(map[string]string)
		fmt.Printf("📢 %s received announcement from %s: %s\n", bu.name, msgData["sender"], msgData["message"])
	case "user_joined":
		username := data.(string)
		fmt.Printf("👋 %s sees that %s joined the chat\n", bu.name, username)
	case "user_left":
		username := data.(string)
		fmt.Printf("👋 %s sees that %s left the chat\n", bu.name, username)
	case "status_changed":
		statusData := data.(map[string]string)
		fmt.Printf("📊 %s sees that %s changed status to: %s\n", bu.name, statusData["user"], statusData["status"])
	case "user_list_response":
		userList := data.(string)
		fmt.Printf("👥 %s received user list: %s\n", bu.name, userList)
	case "error":
		errorMsg := data.(string)
		fmt.Printf("❌ %s received error: %s\n", bu.name, errorMsg)
	default:
		fmt.Printf("🔔 %s received unknown notification: %s\n", bu.name, event)
	}
}