package observers

import (
	"fmt"
	"time"
)

// Observer interface defines the contract for all observers
type Observer interface {
	Update(news *NewsItem)
	GetID() string
	GetType() string
}

// Subject interface defines the contract for subjects being observed
type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	NotifyObservers(news *NewsItem)
	GetSubscriberCount() int
}

// NewsItem represents a news article with metadata
type NewsItem struct {
	ID          string
	Title       string
	Content     string
	Category    string
	Priority    Priority
	PublishedAt time.Time
	Author      string
}

// Priority represents the urgency level of news
type Priority int

const (
	Low Priority = iota
	Medium
	High
	Breaking
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Breaking:
		return "Breaking"
	default:
		return "Unknown"
	}
}

// NewsAgency is the concrete subject that publishes news
type NewsAgency struct {
	name        string
	observers   []Observer
	newsHistory []*NewsItem
}

// NewNewsAgency creates a new news agency
func NewNewsAgency(name string) *NewsAgency {
	return &NewsAgency{
		name:        name,
		observers:   make([]Observer, 0),
		newsHistory: make([]*NewsItem, 0),
	}
}

// Subscribe adds an observer to the notification list
func (na *NewsAgency) Subscribe(observer Observer) {
	na.observers = append(na.observers, observer)
	fmt.Printf("📢 %s (%s) subscribed to %s\n", observer.GetID(), observer.GetType(), na.name)
}

// Unsubscribe removes an observer from the notification list
func (na *NewsAgency) Unsubscribe(observer Observer) {
	for i, obs := range na.observers {
		if obs.GetID() == observer.GetID() {
			// Remove observer from slice
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			fmt.Printf("📤 %s (%s) unsubscribed from %s\n", observer.GetID(), observer.GetType(), na.name)
			return
		}
	}
	fmt.Printf("⚠️  Observer %s not found in subscription list\n", observer.GetID())
}

// NotifyObservers sends news updates to all subscribed observers
func (na *NewsAgency) NotifyObservers(news *NewsItem) {
	fmt.Printf("\n📡 %s broadcasting news to %d subscribers...\n", na.name, len(na.observers))
	fmt.Printf("📰 News: %s [%s Priority]\n", news.Title, news.Priority.String())
	
	for _, observer := range na.observers {
		observer.Update(news)
	}
	
	fmt.Printf("✅ Broadcast complete!\n\n")
}

// PublishNews creates and publishes a new news item
func (na *NewsAgency) PublishNews(title, content, category, author string, priority Priority) {
	news := &NewsItem{
		ID:          fmt.Sprintf("news_%d", time.Now().UnixNano()),
		Title:       title,
		Content:     content,
		Category:    category,
		Priority:    priority,
		PublishedAt: time.Now(),
		Author:      author,
	}
	
	// Store in history
	na.newsHistory = append(na.newsHistory, news)
	
	// Notify all observers
	na.NotifyObservers(news)
}

// GetSubscriberCount returns the number of current subscribers
func (na *NewsAgency) GetSubscriberCount() int {
	return len(na.observers)
}

// GetNewsHistory returns all published news items
func (na *NewsAgency) GetNewsHistory() []*NewsItem {
	return na.newsHistory
}

// GetName returns the agency name
func (na *NewsAgency) GetName() string {
	return na.name
}//
 EmailSubscriber represents a subscriber who receives news via email
type EmailSubscriber struct {
	id           string
	email        string
	name         string
	preferences  []string // Categories of interest
	isActive     bool
	notifications int
}

// NewEmailSubscriber creates a new email subscriber
func NewEmailSubscriber(id, email, name string, preferences []string) *EmailSubscriber {
	return &EmailSubscriber{
		id:          id,
		email:       email,
		name:        name,
		preferences: preferences,
		isActive:    true,
		notifications: 0,
	}
}

func (es *EmailSubscriber) Update(news *NewsItem) {
	if !es.isActive {
		return
	}
	
	// Check if subscriber is interested in this category
	if !es.isInterestedIn(news.Category) {
		return
	}
	
	es.notifications++
	
	fmt.Printf("📧 EMAIL to %s (%s):\n", es.name, es.email)
	fmt.Printf("   Subject: [%s] %s\n", news.Priority.String(), news.Title)
	if news.Priority == Breaking {
		fmt.Printf("   🚨 BREAKING NEWS ALERT! 🚨\n")
	}
	fmt.Printf("   Category: %s | Author: %s\n", news.Category, news.Author)
	fmt.Printf("   Content: %s\n", es.truncateContent(news.Content, 100))
}

func (es *EmailSubscriber) GetID() string {
	return es.id
}

func (es *EmailSubscriber) GetType() string {
	return "Email Subscriber"
}

func (es *EmailSubscriber) isInterestedIn(category string) bool {
	if len(es.preferences) == 0 {
		return true // No preferences means interested in all
	}
	
	for _, pref := range es.preferences {
		if pref == category || pref == "all" {
			return true
		}
	}
	return false
}

func (es *EmailSubscriber) truncateContent(content string, maxLen int) string {
	if len(content) <= maxLen {
		return content
	}
	return content[:maxLen] + "..."
}

func (es *EmailSubscriber) SetActive(active bool) {
	es.isActive = active
}

func (es *EmailSubscriber) GetNotificationCount() int {
	return es.notifications
}

// SMSSubscriber represents a subscriber who receives news via SMS
type SMSSubscriber struct {
	id           string
	phoneNumber  string
	name         string
	onlyBreaking bool // Only receive breaking news
	notifications int
}

// NewSMSSubscriber creates a new SMS subscriber
func NewSMSSubscriber(id, phoneNumber, name string, onlyBreaking bool) *SMSSubscriber {
	return &SMSSubscriber{
		id:           id,
		phoneNumber:  phoneNumber,
		name:         name,
		onlyBreaking: onlyBreaking,
		notifications: 0,
	}
}

func (ss *SMSSubscriber) Update(news *NewsItem) {
	// SMS subscribers might only want breaking news to avoid spam
	if ss.onlyBreaking && news.Priority != Breaking {
		return
	}
	
	ss.notifications++
	
	fmt.Printf("📱 SMS to %s (%s):\n", ss.name, ss.phoneNumber)
	if news.Priority == Breaking {
		fmt.Printf("   🚨 BREAKING: %s\n", news.Title)
	} else {
		fmt.Printf("   📰 %s: %s\n", news.Category, news.Title)
	}
	fmt.Printf("   %s\n", ss.truncateContent(news.Content, 160)) // SMS character limit
}

func (ss *SMSSubscriber) GetID() string {
	return ss.id
}

func (ss *SMSSubscriber) GetType() string {
	return "SMS Subscriber"
}

func (ss *SMSSubscriber) truncateContent(content string, maxLen int) string {
	if len(content) <= maxLen {
		return content
	}
	return content[:maxLen] + "..."
}

func (ss *SMSSubscriber) GetNotificationCount() int {
	return ss.notifications
}

// MobileAppSubscriber represents a subscriber using a mobile app
type MobileAppSubscriber struct {
	id            string
	deviceID      string
	username      string
	pushEnabled   bool
	categories    []string
	notifications int
}

// NewMobileAppSubscriber creates a new mobile app subscriber
func NewMobileAppSubscriber(id, deviceID, username string, categories []string) *MobileAppSubscriber {
	return &MobileAppSubscriber{
		id:          id,
		deviceID:    deviceID,
		username:    username,
		pushEnabled: true,
		categories:  categories,
		notifications: 0,
	}
}

func (mas *MobileAppSubscriber) Update(news *NewsItem) {
	if !mas.pushEnabled {
		return
	}
	
	// Check category preferences
	if !mas.isInterestedIn(news.Category) {
		return
	}
	
	mas.notifications++
	
	fmt.Printf("📲 PUSH NOTIFICATION to %s (Device: %s):\n", mas.username, mas.deviceID[:8]+"...")
	
	// Different notification styles based on priority
	switch news.Priority {
	case Breaking:
		fmt.Printf("   🚨 BREAKING NEWS 🚨\n")
		fmt.Printf("   %s\n", news.Title)
		fmt.Printf("   Tap to read full story\n")
	case High:
		fmt.Printf("   🔴 HIGH PRIORITY\n")
		fmt.Printf("   %s\n", news.Title)
	default:
		fmt.Printf("   📰 %s\n", news.Title)
		fmt.Printf("   %s\n", news.Category)
	}
}

func (mas *MobileAppSubscriber) GetID() string {
	return mas.id
}

func (mas *MobileAppSubscriber) GetType() string {
	return "Mobile App Subscriber"
}

func (mas *MobileAppSubscriber) isInterestedIn(category string) bool {
	if len(mas.categories) == 0 {
		return true
	}
	
	for _, cat := range mas.categories {
		if cat == category || cat == "all" {
			return true
		}
	}
	return false
}

func (mas *MobileAppSubscriber) SetPushEnabled(enabled bool) {
	mas.pushEnabled = enabled
}

func (mas *MobileAppSubscriber) GetNotificationCount() int {
	return mas.notifications
}

// WebSocketSubscriber represents a real-time web subscriber
type WebSocketSubscriber struct {
	id           string
	sessionID    string
	ipAddress    string
	isConnected  bool
	notifications int
}

// NewWebSocketSubscriber creates a new WebSocket subscriber
func NewWebSocketSubscriber(id, sessionID, ipAddress string) *WebSocketSubscriber {
	return &WebSocketSubscriber{
		id:          id,
		sessionID:   sessionID,
		ipAddress:   ipAddress,
		isConnected: true,
		notifications: 0,
	}
}

func (wss *WebSocketSubscriber) Update(news *NewsItem) {
	if !wss.isConnected {
		return
	}
	
	wss.notifications++
	
	fmt.Printf("🌐 WEBSOCKET to Session %s (%s):\n", wss.sessionID[:8]+"...", wss.ipAddress)
	fmt.Printf("   {\n")
	fmt.Printf("     \"type\": \"news_update\",\n")
	fmt.Printf("     \"id\": \"%s\",\n", news.ID)
	fmt.Printf("     \"title\": \"%s\",\n", news.Title)
	fmt.Printf("     \"category\": \"%s\",\n", news.Category)
	fmt.Printf("     \"priority\": \"%s\",\n", news.Priority.String())
	fmt.Printf("     \"timestamp\": \"%s\"\n", news.PublishedAt.Format(time.RFC3339))
	fmt.Printf("   }\n")
}

func (wss *WebSocketSubscriber) GetID() string {
	return wss.id
}

func (wss *WebSocketSubscriber) GetType() string {
	return "WebSocket Subscriber"
}

func (wss *WebSocketSubscriber) Disconnect() {
	wss.isConnected = false
}

func (wss *WebSocketSubscriber) Connect() {
	wss.isConnected = true
}

func (wss *WebSocketSubscriber) GetNotificationCount() int {
	return wss.notifications
}