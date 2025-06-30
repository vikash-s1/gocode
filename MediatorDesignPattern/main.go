package main

import (
	"fmt"
	"mediator/internal/components"
	"mediator/internal/mediator"
	"time"
)

func main() {
	fmt.Println("=== Mediator Design Pattern Demo ===")
	fmt.Println("Simulating a Chat Room Application\n")

	// Create the chat mediator
	chatRoom := mediator.NewChatMediator()

	// Create different types of users
	alice := components.NewRegularUser("Alice")
	bob := components.NewRegularUser("Bob")
	charlie := components.NewModeratorUser("Charlie")
	chatBot := components.NewBotUser("ChatBot")

	fmt.Println("=== Users Joining Chat Room ===")
	// Register users with the mediator
	chatRoom.RegisterComponent(alice)
	time.Sleep(500 * time.Millisecond)
	
	chatRoom.RegisterComponent(bob)
	time.Sleep(500 * time.Millisecond)
	
	chatRoom.RegisterComponent(charlie)
	time.Sleep(500 * time.Millisecond)
	
	chatRoom.RegisterComponent(chatBot)
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Regular Chat Messages ===")
	alice.SendMessage("Hello everyone!")
	time.Sleep(500 * time.Millisecond)
	
	bob.SendMessage("Hey Alice! How's it going?")
	time.Sleep(500 * time.Millisecond)
	
	alice.SendMessage("Great! Nice to meet you all.")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Bot Interaction ===")
	bob.SendMessage("bot hello")
	time.Sleep(500 * time.Millisecond)
	
	alice.SendMessage("ChatBot help")
	time.Sleep(500 * time.Millisecond)
	
	bob.SendMessage("bot joke")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Private Messages ===")
	alice.SendPrivateMessage("Bob", "Want to grab coffee later?")
	time.Sleep(500 * time.Millisecond)
	
	bob.SendPrivateMessage("Alice", "Sure! That sounds great.")
	time.Sleep(500 * time.Millisecond)
	
	// Try to send private message to non-existent user
	alice.SendPrivateMessage("NonExistentUser", "This should fail")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Moderator Actions ===")
	charlie.SendAnnouncement("Welcome to our daily standup meeting!")
	time.Sleep(500 * time.Millisecond)
	
	charlie.SendMessage("Please keep the discussion professional.")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Status Changes ===")
	alice.SetStatus("away")
	time.Sleep(500 * time.Millisecond)
	
	bob.SetStatus("busy")
	time.Sleep(500 * time.Millisecond)
	
	charlie.SetStatus("available")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== User List Request ===")
	alice.RequestUserList()
	time.Sleep(500 * time.Millisecond)
	
	bob.SendMessage("bot users")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== User Leaving Chat ===")
	chatRoom.UnregisterComponent(bob)
	time.Sleep(500 * time.Millisecond)
	
	alice.SendMessage("Where did Bob go?")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Chat Log Summary ===")
	chatLog := chatRoom.GetChatLog()
	fmt.Printf("Total messages logged: %d\n", len(chatLog))
	fmt.Println("Recent chat history:")
	
	// Show last 10 messages
	start := len(chatLog) - 10
	if start < 0 {
		start = 0
	}
	
	for i := start; i < len(chatLog); i++ {
		fmt.Printf("  %s\n", chatLog[i])
	}

	fmt.Println("\n=== Active Users ===")
	activeUsers := chatRoom.GetActiveUsers()
	fmt.Printf("Currently active users: %v\n", activeUsers)

	fmt.Println("\n=== Demo Complete ===")
	fmt.Println("The Mediator pattern successfully coordinated communication")
	fmt.Println("between different types of components without them knowing")
	fmt.Println("about each other directly.")
}