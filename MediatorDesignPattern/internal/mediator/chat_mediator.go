package mediator

import (
	"fmt"
	"strings"
	"time"
)

// ChatMediator implements the Mediator interface for a chat room
type ChatMediator struct {
	components []Component
	chatLog    []string
}

// NewChatMediator creates a new ChatMediator
func NewChatMediator() *ChatMediator {
	return &ChatMediator{
		components: make([]Component, 0),
		chatLog:    make([]string, 0),
	}
}

// RegisterComponent adds a component to the mediator
func (cm *ChatMediator) RegisterComponent(component Component) {
	cm.components = append(cm.components, component)
	component.SetMediator(cm)
	
	// Notify all other components about new user joining
	joinMessage := fmt.Sprintf("%s joined the chat", component.GetName())
	cm.logMessage("SYSTEM", joinMessage)
	
	for _, comp := range cm.components {
		if comp != component {
			comp.HandleNotification("user_joined", component.GetName())
		}
	}
}

// UnregisterComponent removes a component from the mediator
func (cm *ChatMediator) UnregisterComponent(component Component) {
	for i, comp := range cm.components {
		if comp == component {
			cm.components = append(cm.components[:i], cm.components[i+1:]...)
			break
		}
	}
	
	// Notify all remaining components about user leaving
	leaveMessage := fmt.Sprintf("%s left the chat", component.GetName())
	cm.logMessage("SYSTEM", leaveMessage)
	
	for _, comp := range cm.components {
		comp.HandleNotification("user_left", component.GetName())
	}
}

// Notify handles communication between components
func (cm *ChatMediator) Notify(sender Component, event string, data interface{}) {
	switch event {
	case "send_message":
		cm.handleMessage(sender, data.(string))
	case "send_private_message":
		cm.handlePrivateMessage(sender, data.(map[string]string))
	case "broadcast_announcement":
		cm.handleAnnouncement(sender, data.(string))
	case "request_user_list":
		cm.handleUserListRequest(sender)
	case "change_status":
		cm.handleStatusChange(sender, data.(string))
	default:
		fmt.Printf("Unknown event: %s\n", event)
	}
}

// handleMessage broadcasts a message to all components except sender
func (cm *ChatMediator) handleMessage(sender Component, message string) {
	logEntry := fmt.Sprintf("%s: %s", sender.GetName(), message)
	cm.logMessage(sender.GetName(), message)
	
	for _, component := range cm.components {
		if component != sender {
			component.HandleNotification("message_received", map[string]string{
				"sender":  sender.GetName(),
				"message": message,
			})
		}
	}
}

// handlePrivateMessage sends a private message to a specific user
func (cm *ChatMediator) handlePrivateMessage(sender Component, data map[string]string) {
	recipient := data["recipient"]
	message := data["message"]
	
	logEntry := fmt.Sprintf("%s -> %s (private): %s", sender.GetName(), recipient, message)
	cm.logMessage("PRIVATE", logEntry)
	
	for _, component := range cm.components {
		if component.GetName() == recipient {
			component.HandleNotification("private_message_received", map[string]string{
				"sender":  sender.GetName(),
				"message": message,
			})
			return
		}
	}
	
	// Notify sender if recipient not found
	sender.HandleNotification("error", "User not found: "+recipient)
}

// handleAnnouncement broadcasts an announcement to all users
func (cm *ChatMediator) handleAnnouncement(sender Component, message string) {
	logEntry := fmt.Sprintf("📢 ANNOUNCEMENT from %s: %s", sender.GetName(), message)
	cm.logMessage("ANNOUNCEMENT", logEntry)
	
	for _, component := range cm.components {
		if component != sender {
			component.HandleNotification("announcement_received", map[string]string{
				"sender":  sender.GetName(),
				"message": message,
			})
		}
	}
}

// handleUserListRequest sends the list of active users to the requester
func (cm *ChatMediator) handleUserListRequest(sender Component) {
	userList := make([]string, 0)
	for _, component := range cm.components {
		userList = append(userList, component.GetName())
	}
	
	sender.HandleNotification("user_list_response", strings.Join(userList, ", "))
}

// handleStatusChange notifies all users about status change
func (cm *ChatMediator) handleStatusChange(sender Component, status string) {
	logEntry := fmt.Sprintf("%s changed status to: %s", sender.GetName(), status)
	cm.logMessage("STATUS", logEntry)
	
	for _, component := range cm.components {
		if component != sender {
			component.HandleNotification("status_changed", map[string]string{
				"user":   sender.GetName(),
				"status": status,
			})
		}
	}
}

// logMessage adds a message to the chat log with timestamp
func (cm *ChatMediator) logMessage(sender, message string) {
	timestamp := time.Now().Format("15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s", timestamp, sender, message)
	cm.chatLog = append(cm.chatLog, logEntry)
}

// GetChatLog returns the complete chat log
func (cm *ChatMediator) GetChatLog() []string {
	return cm.chatLog
}

// GetActiveUsers returns the list of active users
func (cm *ChatMediator) GetActiveUsers() []string {
	users := make([]string, 0)
	for _, component := range cm.components {
		users = append(users, component.GetName())
	}
	return users
}