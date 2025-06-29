package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"observerpattern/internal/observers"
)

func main() {
	fmt.Println("📰 Welcome to the Observer Pattern News Agency Demo!")
	fmt.Println(strings.Repeat("=", 55))

	// Interactive demo
	runInteractiveDemo()

	fmt.Println("\n🎯 Automated Demo:")
	fmt.Println(strings.Repeat("=", 30))

	// Automated demo to show all observer patterns
	runAutomatedDemo()
}

func runInteractiveDemo() {
	scanner := bufio.NewScanner(os.Stdin)

	// Create news agency
	newsAgency := observers.NewNewsAgency("Global News Network")

	// Create some initial subscribers
	setupInitialSubscribers(newsAgency)

	fmt.Printf("\n🏢 %s is ready!\n", newsAgency.GetName())
	fmt.Printf("📊 Current subscribers: %d\n", newsAgency.GetSubscriberCount())

	fmt.Println("\nCommands:")
	fmt.Println("1 - Publish Breaking News")
	fmt.Println("2 - Publish Regular News")
	fmt.Println("3 - Add New Subscriber")
	fmt.Println("4 - Remove Subscriber")
	fmt.Println("5 - Show Subscriber Stats")
	fmt.Println("6 - Show News History")
	fmt.Println("q - Quit to Automated Demo")

	for {
		fmt.Printf("\nSubscribers: %d | Enter command: ", newsAgency.GetSubscriberCount())
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			publishBreakingNews(newsAgency, scanner)
		case "2":
			publishRegularNews(newsAgency, scanner)
		case "3":
			addNewSubscriber(newsAgency, scanner)
		case "4":
			removeSubscriber(newsAgency, scanner)
		case "5":
			showSubscriberStats(newsAgency)
		case "6":
			showNewsHistory(newsAgency)
		case "q":
			return
		default:
			fmt.Println("❌ Invalid command. Try again.")
		}
	}
}

func setupInitialSubscribers(newsAgency *observers.NewsAgency) {
	// Email subscribers with different preferences
	emailSub1 := observers.NewEmailSubscriber("email_001", "john@email.com", "John Doe", []string{"technology", "business"})
	emailSub2 := observers.NewEmailSubscriber("email_002", "jane@email.com", "Jane Smith", []string{"all"})

	// SMS subscribers (only breaking news to avoid spam)
	smsSub1 := observers.NewSMSSubscriber("sms_001", "+1-555-0101", "Bob Johnson", true)
	smsSub2 := observers.NewSMSSubscriber("sms_002", "+1-555-0102", "Alice Brown", false)

	// Mobile app subscribers
	mobileSub1 := observers.NewMobileAppSubscriber("mobile_001", "device_abc123", "TechGuru", []string{"technology", "science"})
	mobileSub2 := observers.NewMobileAppSubscriber("mobile_002", "device_xyz789", "NewsReader", []string{"politics", "world"})

	// WebSocket subscribers
	wsSub1 := observers.NewWebSocketSubscriber("ws_001", "session_12345", "192.168.1.100")

	// Subscribe all to news agency
	newsAgency.Subscribe(emailSub1)
	newsAgency.Subscribe(emailSub2)
	newsAgency.Subscribe(smsSub1)
	newsAgency.Subscribe(smsSub2)
	newsAgency.Subscribe(mobileSub1)
	newsAgency.Subscribe(mobileSub2)
	newsAgency.Subscribe(wsSub1)
}

func publishBreakingNews(newsAgency *observers.NewsAgency, scanner *bufio.Scanner) {
	fmt.Print("Enter breaking news title: ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = "Major Economic Announcement Expected"
	}

	fmt.Print("Enter news content: ")
	scanner.Scan()
	content := scanner.Text()
	if content == "" {
		content = "Government officials are expected to make a significant economic announcement that could impact global markets."
	}

	fmt.Print("Enter category (politics/business/technology/world/sports): ")
	scanner.Scan()
	category := scanner.Text()
	if category == "" {
		category = "business"
	}

	newsAgency.PublishNews(title, content, category, "Breaking News Team", observers.Breaking)
}

func publishRegularNews(newsAgency *observers.NewsAgency, scanner *bufio.Scanner) {
	fmt.Print("Enter news title: ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = "New Technology Breakthrough Announced"
	}

	fmt.Print("Enter news content: ")
	scanner.Scan()
	content := scanner.Text()
	if content == "" {
		content = "Researchers have developed a new technology that promises to revolutionize the industry."
	}

	fmt.Print("Enter category: ")
	scanner.Scan()
	category := scanner.Text()
	if category == "" {
		category = "technology"
	}

	fmt.Print("Enter priority (0=Low, 1=Medium, 2=High): ")
	scanner.Scan()
	priorityStr := scanner.Text()
	priority := observers.Medium
	if p, err := strconv.Atoi(priorityStr); err == nil && p >= 0 && p <= 2 {
		priority = observers.Priority(p)
	}

	newsAgency.PublishNews(title, content, category, "News Team", priority)
}

func addNewSubscriber(newsAgency *observers.NewsAgency, scanner *bufio.Scanner) {
	fmt.Println("Choose subscriber type:")
	fmt.Println("1 - Email Subscriber")
	fmt.Println("2 - SMS Subscriber")
	fmt.Println("3 - Mobile App Subscriber")
	fmt.Println("4 - WebSocket Subscriber")

	fmt.Print("Enter choice: ")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		fmt.Print("Enter email: ")
		scanner.Scan()
		email := scanner.Text()
		fmt.Print("Enter name: ")
		scanner.Scan()
		name := scanner.Text()
		
		id := fmt.Sprintf("email_%d", time.Now().UnixNano())
		subscriber := observers.NewEmailSubscriber(id, email, name, []string{"all"})
		newsAgency.Subscribe(subscriber)

	case "2":
		fmt.Print("Enter phone number: ")
		scanner.Scan()
		phone := scanner.Text()
		fmt.Print("Enter name: ")
		scanner.Scan()
		name := scanner.Text()
		
		id := fmt.Sprintf("sms_%d", time.Now().UnixNano())
		subscriber := observers.NewSMSSubscriber(id, phone, name, true)
		newsAgency.Subscribe(subscriber)

	case "3":
		fmt.Print("Enter username: ")
		scanner.Scan()
		username := scanner.Text()
		
		id := fmt.Sprintf("mobile_%d", time.Now().UnixNano())
		deviceID := fmt.Sprintf("device_%d", time.Now().UnixNano())
		subscriber := observers.NewMobileAppSubscriber(id, deviceID, username, []string{"all"})
		newsAgency.Subscribe(subscriber)

	case "4":
		fmt.Print("Enter IP address: ")
		scanner.Scan()
		ip := scanner.Text()
		
		id := fmt.Sprintf("ws_%d", time.Now().UnixNano())
		sessionID := fmt.Sprintf("session_%d", time.Now().UnixNano())
		subscriber := observers.NewWebSocketSubscriber(id, sessionID, ip)
		newsAgency.Subscribe(subscriber)

	default:
		fmt.Println("❌ Invalid choice")
	}
}

func removeSubscriber(newsAgency *observers.NewsAgency, scanner *bufio.Scanner) {
	fmt.Print("Enter subscriber ID to remove: ")
	scanner.Scan()
	id := scanner.Text()
	
	// For demo purposes, we'll create a dummy observer with the given ID
	// In a real system, you'd maintain a registry of observers
	fmt.Printf("⚠️  Note: In this demo, we can't remove specific subscribers by ID.\n")
	fmt.Printf("In a real system, you'd maintain a registry of observers for removal.\n")
}

func showSubscriberStats(newsAgency *observers.NewsAgency) {
	fmt.Printf("\n📊 Subscriber Statistics:\n")
	fmt.Printf("Total Subscribers: %d\n", newsAgency.GetSubscriberCount())
	fmt.Printf("News Agency: %s\n", newsAgency.GetName())
	fmt.Printf("Total News Published: %d\n", len(newsAgency.GetNewsHistory()))
}

func showNewsHistory(newsAgency *observers.NewsAgency) {
	history := newsAgency.GetNewsHistory()
	fmt.Printf("\n📚 News History (%d items):\n", len(history))
	fmt.Println(strings.Repeat("-", 50))
	
	for i, news := range history {
		fmt.Printf("%d. [%s] %s\n", i+1, news.Priority.String(), news.Title)
		fmt.Printf("   Category: %s | Author: %s\n", news.Category, news.Author)
		fmt.Printf("   Published: %s\n", news.PublishedAt.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}
}

func runAutomatedDemo() {
	// Create news agency
	newsAgency := observers.NewNewsAgency("Tech News Daily")

	// Demo 1: Basic Observer Pattern
	fmt.Println("\n🎬 Demo 1: Basic Observer Pattern")
	fmt.Println(strings.Repeat("-", 40))

	// Create different types of subscribers
	emailSub := observers.NewEmailSubscriber("demo_email", "demo@example.com", "Demo User", []string{"technology"})
	smsSub := observers.NewSMSSubscriber("demo_sms", "+1-555-DEMO", "SMS User", false)
	mobileSub := observers.NewMobileAppSubscriber("demo_mobile", "demo_device_123", "MobileUser", []string{"technology", "business"})
	wsSub := observers.NewWebSocketSubscriber("demo_ws", "demo_session_456", "127.0.0.1")

	// Subscribe all observers
	newsAgency.Subscribe(emailSub)
	newsAgency.Subscribe(smsSub)
	newsAgency.Subscribe(mobileSub)
	newsAgency.Subscribe(wsSub)

	// Publish news and see all observers get notified
	newsAgency.PublishNews(
		"Revolutionary AI Breakthrough",
		"Scientists have developed an AI system that can solve complex problems with unprecedented accuracy.",
		"technology",
		"Tech Reporter",
		observers.High,
	)

	time.Sleep(1 * time.Second)

	// Demo 2: Breaking News Notification
	fmt.Println("\n🎬 Demo 2: Breaking News Alert")
	fmt.Println(strings.Repeat("-", 35))

	newsAgency.PublishNews(
		"Major Security Vulnerability Discovered",
		"A critical security flaw has been found in widely-used software, affecting millions of users worldwide.",
		"technology",
		"Security Team",
		observers.Breaking,
	)

	time.Sleep(1 * time.Second)

	// Demo 3: Observer Unsubscription
	fmt.Println("\n🎬 Demo 3: Observer Unsubscription")
	fmt.Println(strings.Repeat("-", 38))

	fmt.Printf("Before unsubscription: %d subscribers\n", newsAgency.GetSubscriberCount())
	newsAgency.Unsubscribe(smsSub)
	fmt.Printf("After unsubscription: %d subscribers\n", newsAgency.GetSubscriberCount())

	// Publish news to see fewer notifications
	newsAgency.PublishNews(
		"New Programming Language Released",
		"A new programming language designed for modern applications has been officially released.",
		"technology",
		"Dev Team",
		observers.Medium,
	)

	time.Sleep(1 * time.Second)

	// Demo 4: Category-based Filtering
	fmt.Println("\n🎬 Demo 4: Category-based Filtering")
	fmt.Println(strings.Repeat("-", 40))

	// Add subscriber interested only in business news
	businessSub := observers.NewEmailSubscriber("business_sub", "business@example.com", "Business Reader", []string{"business"})
	newsAgency.Subscribe(businessSub)

	// Publish business news
	newsAgency.PublishNews(
		"Stock Market Reaches New High",
		"Major stock indices have reached record highs following positive economic indicators.",
		"business",
		"Financial Reporter",
		observers.High,
	)

	time.Sleep(1 * time.Second)

	// Publish sports news (business subscriber won't be notified)
	newsAgency.PublishNews(
		"Championship Game Tonight",
		"The final championship game is scheduled for tonight with record ticket sales.",
		"sports",
		"Sports Reporter",
		observers.Low,
	)

	// Demo 5: Multiple News Agencies
	fmt.Println("\n🎬 Demo 5: Multiple News Agencies")
	fmt.Println(strings.Repeat("-", 40))

	// Create another news agency
	sportsAgency := observers.NewNewsAgency("Sports Central")
	
	// Same subscriber can subscribe to multiple agencies
	sportsAgency.Subscribe(mobileSub)
	sportsAgency.Subscribe(wsSub)

	sportsAgency.PublishNews(
		"Trade Deadline Approaching",
		"Teams are making final moves before the trade deadline expires at midnight.",
		"sports",
		"Sports Insider",
		observers.Medium,
	)

	fmt.Println("\n✨ Observer Pattern Demo Complete!")
	fmt.Println("Key takeaways:")
	fmt.Println("• Observers are automatically notified when subjects change")
	fmt.Println("• Loose coupling between subjects and observers")
	fmt.Println("• Dynamic subscription/unsubscription")
	fmt.Println("• Different observer types can handle notifications differently")
	fmt.Println("• Filtering and customization per observer")
}