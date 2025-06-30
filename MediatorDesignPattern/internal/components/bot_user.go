package components

import (
	"fmt"
	"strings"
)

// BotUser represents an automated chat bot
type BotUser struct {
	*BaseUser
	responses map[string]string
}

// NewBotUser creates a new BotUser
func NewBotUser(name string) *BotUser {
	bot := &BotUser{
		BaseUser:  NewBaseUser(name),
		responses: make(map[string]string),
	}
	
	// Initialize bot responses
	bot.responses["hello"] = "Hello there! How can I help you today?"
	bot.responses["help"] = "Available commands: hello, help, time, users, joke"
	bot.responses["time"] = "I don't have access to time functions, sorry!"
	bot.responses["users"] = "Let me get the user list for you..."
	bot.responses["joke"] = "Why do programmers prefer dark mode? Because light attracts bugs! 🐛"
	
	return bot
}

// HandleNotification handles notifications with bot-specific behavior
func (bu *BotUser) HandleNotification(event string, data interface{}) {
	switch event {
	case "message_received":
		msgData := data.(map[string]string)
		sender := msgData["sender"]
		message := strings.ToLower(msgData["message"])
		
		fmt.Printf("🤖 Bot %s received message from %s: %s\n", bu.name, sender, msgData["message"])
		
		// Check if message is directed at the bot
		if strings.Contains(message, strings.ToLower(bu.name)) || strings.HasPrefix(message, "bot") {
			bu.respondToMessage(sender, message)
		}
		
	case "user_joined":
		username := data.(string)
		fmt.Printf("🤖 Bot %s welcomes %s\n", bu.name, username)
		welcomeMsg := fmt.Sprintf("Welcome to the chat, %s! Type 'bot help' for available commands.", username)
		bu.SendMessage(welcomeMsg)
		
	default:
		// Handle other notifications normally
		bu.BaseUser.HandleNotification(event, data)
	}
}

// respondToMessage generates appropriate responses to user messages
func (bu *BotUser) respondToMessage(sender, message string) {
	var response string
	found := false
	
	// Check for known commands
	for command, resp := range bu.responses {
		if strings.Contains(message, command) {
			response = resp
			found = true
			break
		}
	}
	
	if !found {
		response = fmt.Sprintf("Sorry %s, I didn't understand that. Type 'bot help' for available commands.", sender)
	}
	
	// Special handling for users command
	if strings.Contains(message, "users") {
		bu.RequestUserList()
		return
	}
	
	// Send response
	bu.SendMessage(response)
}