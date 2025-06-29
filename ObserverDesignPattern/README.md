# Observer Design Pattern in Go

## Overview

The Observer Design Pattern is a behavioral design pattern that defines a one-to-many dependency between objects. When one object (the subject) changes state, all its dependents (observers) are notified and updated automatically. This pattern is also known as the Publish-Subscribe pattern.

## Problem It Solves

Without the Observer pattern, you'd typically have tight coupling between objects that need to stay synchronized:

```go
// BAD: Tight coupling and manual notification management
type NewsAgency struct {
    emailService *EmailService
    smsService   *SMSService
    pushService  *PushService
}

func (na *NewsAgency) PublishNews(news string) {
    // Manually notify each service - tightly coupled!
    na.emailService.SendEmail(news)
    na.smsService.SendSMS(news)
    na.pushService.SendPush(news)
    
    // Adding new notification types requires modifying this code
}
```

This leads to:
- **Tight Coupling**: Subject must know about all its dependents
- **Violation of Open/Closed Principle**: Adding new observers requires modifying existing code
- **Difficult Maintenance**: Changes to notification logic affect multiple classes
- **No Dynamic Subscription**: Can't add/remove observers at runtime

## Solution

The Observer pattern decouples subjects from observers by introducing interfaces and allowing dynamic subscription management.

## Implementation Structure

### Core Components

1. **Observer Interface** (`Observer`)
   - Defines the contract for objects that should be notified
   - Contains `Update()` method called when subject changes

2. **Subject Interface** (`Subject`)
   - Defines methods for managing observers
   - `Subscribe()`, `Unsubscribe()`, `NotifyObservers()`

3. **Concrete Subject** (`NewsAgency`)
   - Implements the Subject interface
   - Maintains list of observers
   - Notifies observers when state changes

4. **Concrete Observers**
   - `EmailSubscriber`: Sends email notifications
   - `SMSSubscriber`: Sends SMS notifications
   - `MobileAppSubscriber`: Sends push notifications
   - `WebSocketSubscriber`: Sends real-time web updates

## News Agency Example

Our implementation simulates a news agency system where different types of subscribers get notified of breaking news:

```
┌─────────────────┐
│   NewsAgency    │
│   (Subject)     │
│                 │
│ - observers[]   │
│ + Subscribe()   │
│ + Unsubscribe() │
│ + NotifyObservers()│
│ + PublishNews() │
└─────────────────┘
         │
         │ notifies
         ▼
┌─────────────────┐
│    Observer     │
│   (Interface)   │
│                 │
│ + Update()      │
│ + GetID()       │
│ + GetType()     │
└─────────────────┘
         ▲
    ┌────┼────┐
    │    │    │
┌───────────┐ ┌──────────────┐ ┌─────────────────┐
│EmailSub   │ │SMSSubscriber │ │MobileAppSub     │
│           │ │              │ │                 │
│+ Update() │ │+ Update()    │ │+ Update()       │
└───────────┘ └──────────────┘ └─────────────────┘
```

## Step-by-Step Implementation

### Step 1: Define Observer Interface

```go
type Observer interface {
    Update(news *NewsItem)
    GetID() string
    GetType() string
}
```

**Purpose**: Establishes the contract that all observers must implement, ensuring they can receive and handle notifications.

### Step 2: Define Subject Interface

```go
type Subject interface {
    Subscribe(observer Observer)
    Unsubscribe(observer Observer)
    NotifyObservers(news *NewsItem)
    GetSubscriberCount() int
}
```

**Purpose**: Defines the contract for objects that can be observed, providing methods for observer management.

### Step 3: Implement Concrete Subject

```go
type NewsAgency struct {
    name        string
    observers   []Observer
    newsHistory []*NewsItem
}

func (na *NewsAgency) Subscribe(observer Observer) {
    na.observers = append(na.observers, observer)
}

func (na *NewsAgency) NotifyObservers(news *NewsItem) {
    for _, observer := range na.observers {
        observer.Update(news)
    }
}
```

**Key Features**:
- Maintains list of observers
- Provides subscription management
- Broadcasts notifications to all observers

### Step 4: Implement Concrete Observers

#### EmailSubscriber
```go
type EmailSubscriber struct {
    id          string
    email       string
    name        string
    preferences []string // Category filtering
}

func (es *EmailSubscriber) Update(news *NewsItem) {
    if !es.isInterestedIn(news.Category) {
        return // Skip if not interested in this category
    }
    
    fmt.Printf("📧 EMAIL to %s: %s\n", es.name, news.Title)
    // Send email notification...
}
```

**Characteristics**:
- Category-based filtering
- Rich email formatting
- Preference management

#### SMSSubscriber
```go
type SMSSubscriber struct {
    id           string
    phoneNumber  string
    onlyBreaking bool // Only breaking news to avoid spam
}

func (ss *SMSSubscriber) Update(news *NewsItem) {
    if ss.onlyBreaking && news.Priority != Breaking {
        return // Skip non-breaking news
    }
    
    fmt.Printf("📱 SMS to %s: %s\n", ss.phoneNumber, news.Title)
    // Send SMS with character limit...
}
```

**Characteristics**:
- Breaking news filtering
- Character limit handling
- Spam prevention

## Key Benefits Demonstrated

### 1. **Loose Coupling**
```go
// Subject doesn't need to know specific observer types
func (na *NewsAgency) NotifyObservers(news *NewsItem) {
    for _, observer := range na.observers {
        observer.Update(news) // Polymorphic call
    }
}
```

### 2. **Dynamic Subscription Management**
```go
// Add observers at runtime
newsAgency.Subscribe(newEmailSubscriber)

// Remove observers when no longer needed
newsAgency.Unsubscribe(oldSubscriber)
```

### 3. **Open/Closed Principle**
```go
// Easy to add new observer types without modifying existing code
type SlackSubscriber struct {
    channelID string
    botToken  string
}

func (ss *SlackSubscriber) Update(news *NewsItem) {
    // Send Slack notification
}
```

### 4. **Broadcast Communication**
```go
// One call notifies all subscribers
newsAgency.PublishNews("Breaking News", content, "politics", "Reporter", Breaking)
// Automatically notifies: Email, SMS, Mobile App, WebSocket subscribers
```

## Running the Example

### Build and Run
```bash
cd ObserverDesignPattern
go run main.go
```

### Interactive Mode Features
1. **Publish Breaking News**: Send high-priority notifications to all subscribers
2. **Publish Regular News**: Send normal priority news with category filtering
3. **Add New Subscriber**: Dynamically add different types of subscribers
4. **Remove Subscriber**: Unsubscribe observers from notifications
5. **Show Statistics**: View subscriber counts and news history
6. **News History**: Review all published news items

### Automated Demos
The program demonstrates:
1. **Basic Observer Pattern**: Multiple observer types receiving notifications
2. **Breaking News Alerts**: High-priority notifications with special formatting
3. **Dynamic Unsubscription**: Removing observers and seeing fewer notifications
4. **Category Filtering**: Observers receiving only relevant news categories
5. **Multiple Subjects**: Same observer subscribing to different news agencies

## Observer Types Comparison

| Observer Type | Notification Speed | Content Length | Filtering | Special Features |
|---------------|-------------------|----------------|-----------|------------------|
| Email | Medium | Unlimited | Category-based | Rich formatting, attachments |
| SMS | Fast | 160 characters | Breaking news only | Character limits, spam prevention |
| Mobile App | Fast | Medium | Category-based | Push notifications, priority levels |
| WebSocket | Instant | Unlimited | None | Real-time, JSON format |

## When to Use Observer Pattern

✅ **Use When**:
- Multiple objects need to be notified of state changes
- You want loose coupling between subjects and observers
- The number of observers varies or changes at runtime
- You need broadcast communication (one-to-many)
- You want to follow the "Don't call us, we'll call you" principle

❌ **Avoid When**:
- Simple one-to-one relationships
- Performance is critical (observer notification has overhead)
- Observer order matters (observers are notified in arbitrary order)
- Complex observer dependencies exist

## Real-World Applications

### Model-View-Controller (MVC)
```go
// Model notifies Views when data changes
type UserModel struct {
    observers []Observer
    userData  *User
}

func (um *UserModel) UpdateUser(user *User) {
    um.userData = user
    um.NotifyObservers(user) // Views update automatically
}
```

### Event Systems
```go
// UI events notify multiple handlers
type Button struct {
    clickHandlers []ClickHandler
}

func (b *Button) OnClick() {
    for _, handler := range b.clickHandlers {
        handler.HandleClick()
    }
}
```

### Stock Market Systems
```go
// Stock price changes notify multiple displays
type Stock struct {
    symbol    string
    price     float64
    observers []StockObserver
}

func (s *Stock) SetPrice(price float64) {
    s.price = price
    s.NotifyObservers() // Charts, alerts, portfolios update
}
```

### Logging Systems
```go
// Log events sent to multiple destinations
type Logger struct {
    handlers []LogHandler
}

func (l *Logger) Log(message string) {
    for _, handler := range l.handlers {
        handler.Handle(message) // File, console, network, etc.
    }
}
```

## Comparison with Other Patterns

### vs Mediator Pattern
- **Observer**: One-to-many communication, subjects broadcast to observers
- **Mediator**: Many-to-many communication through central mediator

### vs Command Pattern
- **Observer**: Notification-based, observers react to state changes
- **Command**: Request-based, commands encapsulate actions

### vs Publish-Subscribe Pattern
- **Observer**: Direct reference between subject and observers
- **Pub-Sub**: Indirect communication through message broker/event bus

## Advanced Considerations

### Thread Safety
For concurrent environments:
```go
import "sync"

type ThreadSafeNewsAgency struct {
    mu        sync.RWMutex
    observers []Observer
}

func (na *ThreadSafeNewsAgency) Subscribe(observer Observer) {
    na.mu.Lock()
    defer na.mu.Unlock()
    na.observers = append(na.observers, observer)
}

func (na *ThreadSafeNewsAgency) NotifyObservers(news *NewsItem) {
    na.mu.RLock()
    observers := make([]Observer, len(na.observers))
    copy(observers, na.observers)
    na.mu.RUnlock()
    
    for _, observer := range observers {
        observer.Update(news)
    }
}
```

### Error Handling
Robust observer notification:
```go
func (na *NewsAgency) NotifyObservers(news *NewsItem) {
    for _, observer := range na.observers {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    log.Printf("Observer %s panicked: %v", observer.GetID(), r)
                }
            }()
            observer.Update(news)
        }()
    }
}
```

### Asynchronous Notifications
For non-blocking notifications:
```go
func (na *NewsAgency) NotifyObserversAsync(news *NewsItem) {
    for _, observer := range na.observers {
        go func(obs Observer) {
            obs.Update(news)
        }(observer)
    }
}
```

### Observer Priority
For ordered notifications:
```go
type PriorityObserver struct {
    Observer
    Priority int
}

func (na *NewsAgency) NotifyObserversByPriority(news *NewsItem) {
    // Sort observers by priority
    sort.Slice(na.priorityObservers, func(i, j int) bool {
        return na.priorityObservers[i].Priority > na.priorityObservers[j].Priority
    })
    
    for _, observer := range na.priorityObservers {
        observer.Update(news)
    }
}
```

### Memory Leak Prevention
Weak references to prevent memory leaks:
```go
type WeakObserver struct {
    observer Observer
    isValid  func() bool
}

func (na *NewsAgency) NotifyObservers(news *NewsItem) {
    validObservers := make([]Observer, 0)
    
    for _, weakObs := range na.weakObservers {
        if weakObs.isValid() {
            validObservers = append(validObservers, weakObs.observer)
            weakObs.observer.Update(news)
        }
    }
    
    // Update list to remove invalid observers
    na.updateObserverList(validObservers)
}
```

This implementation demonstrates how the Observer pattern creates flexible, maintainable notification systems by establishing loose coupling between subjects and observers, enabling dynamic subscription management, and supporting broadcast communication patterns.