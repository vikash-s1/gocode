# Adapter Design Pattern in Go

This project demonstrates the **Adapter Design Pattern** implementation in Go using a private package structure. The Adapter pattern allows incompatible interfaces to work together by creating a bridge between them.

## 🎯 What is the Adapter Pattern?

The Adapter pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It acts as a wrapper between two objects, catching calls for one object and transforming them to format and interface recognizable by the second object.

### Real-World Analogy
Think of a power adapter when traveling internationally. Your laptop charger (client) expects a specific plug format, but the wall outlet (service) in another country has a different format. The power adapter (adapter) allows your charger to work with the foreign outlet.

## 🏗️ Project Structure

```
adapter-pattern/
├── main.go                           # Main application demonstrating the pattern
├── go.mod                           # Go module definition
├── README.md                        # This file
└── internal/                        # Private packages
    ├── adapters/                    # Adapter implementations
    │   ├── payment_adapters.go      # Payment system adapters
    │   ├── database_adapters.go     # Database system adapters
    │   └── media_adapters.go        # Media player adapters
    ├── legacy/                      # Legacy systems with incompatible interfaces
    │   ├── payment_systems.go       # Old payment gateways
    │   ├── database_systems.go      # Old database systems
    │   └── media_players.go         # Old media players
    ├── modern/                      # Modern target interfaces
    │   └── interfaces.go            # Standard interfaces we want to use
    └── payment/                     # Payment-specific interfaces
        └── interfaces.go            # Payment processor interface
```

## 🔧 Key Components

### 1. Target Interface (What we want)
Modern, standardized interfaces that our application expects:
- `payment.Processor` - Unified payment processing interface
- `modern.Database` - Unified database interface  
- `modern.MediaPlayer` - Unified media player interface

### 2. Adaptee (What we have)
Legacy systems with their own incompatible interfaces:
- `legacy.PayPalGateway` - PayPal's specific API
- `legacy.StripeGateway` - Stripe's specific API
- `legacy.MySQLDatabase` - MySQL's specific methods
- `legacy.MP3Player` - MP3 player's specific methods

### 3. Adapter (The Bridge)
Adapters that implement the target interface while wrapping the adaptee:
- `PayPalAdapter` - Adapts PayPal to `payment.Processor`
- `MySQLAdapter` - Adapts MySQL to `modern.Database`
- `MP3Adapter` - Adapts MP3Player to `modern.MediaPlayer`

## 🚀 Running the Example

### Prerequisites
- Go 1.21 or later
- No external dependencies required

### Installation & Execution

```bash
# Clone or download the project
git clone <repository-url>
cd adapter-pattern

# Initialize Go module (if needed)
go mod init github.com/example/adapter-pattern

# Run the demonstration
go run main.go
```

### Expected Output

```
=== Adapter Design Pattern Demo ===

🏦 Payment Processing Adapters
------------------------------

1. Processing $99.99 payment:
   ✓ PayPal: Processing $99.99 USD (Transaction: PP_1640995200)
   Status: PAYPAL_ACTIVE

2. Processing $99.99 payment:
   ✓ Stripe: Charged 9999 cents (Charge: ch_1640995200)
   Status: STRIPE_SUCCESS

3. Processing $99.99 payment:
   ✓ Bitcoin: Sent 0.00222200 BTC (TX: btc_61c8a000)
   Status: BTC_NETWORK_ACTIVE (Confirmations: 0)

==================================================

🗄️  Database Connection Adapters
--------------------------------

1. Testing database connection:
   ✓ MySQL: Opening connection to mysql.example.com:3306
   ✓ MySQL: Executing SQL: SELECT * FROM users LIMIT 5
   ✓ MySQL: Closing connection

2. Testing database connection:
   ✓ PostgreSQL: Establishing connection with: postgres://user:pass@localhost/db
   ✓ PostgreSQL: Executing query: SELECT * FROM users LIMIT 5
   ✓ PostgreSQL: Terminating connection

3. Testing database connection:
   ✓ MongoDB: Starting session with URI: mongodb://localhost:27017
   ✓ MongoDB: Finding documents with filter: {}
   ✓ MongoDB: Ending session

==================================================

🎵 Media Player Adapters
------------------------

1. Playing song1.mp3:
   ♪ MP3Player: Playing song1.mp3
   🔊 MP3Player: Volume set to 75%
   Volume set to: 75%
   ⏹ MP3Player: Stopped playing song1.mp3

2. Playing song2.wav:
   ♪ WAVPlayer: Starting playback of song2.wav
   🔊 WAVPlayer: Audio level adjusted to 75%
   Volume set to: 75%
   ⏹ WAVPlayer: Halted playback of song2.wav

3. Playing song3.flac:
   ♪ FLACPlayer: Loading and playing song3.flac
   🔊 FLACPlayer: Sound level modified to 75%
   Volume set to: 75%
   ⏹ FLACPlayer: Terminated playback of song3.flac
```

## 💡 Pattern Benefits Demonstrated

### 1. **Interface Unification**
```go
// Before: Different interfaces for each payment system
paypal.MakePayment(99.99, "USD")
stripe.ChargeCard(9999, "card_token")
bitcoin.SendBitcoin(0.002222, "address")

// After: Unified interface through adapters
for _, processor := range paymentProcessors {
    processor.ProcessPayment(99.99)  // Same method for all!
}
```

### 2. **Legacy System Integration**
The pattern allows you to integrate legacy systems without modifying their code:
```go
// Legacy system remains unchanged
type PayPalGateway struct { /* existing code */ }
func (p *PayPalGateway) MakePayment(dollars float64, currency string) (string, error)

// Adapter bridges the gap
type PayPalAdapter struct {
    gateway *legacy.PayPalGateway
}
func (p *PayPalAdapter) ProcessPayment(amount float64) error {
    _, err := p.gateway.MakePayment(amount, "USD")
    return err
}
```

### 3. **Polymorphic Usage**
All adapted systems can be used interchangeably:
```go
func processPayments(processors []payment.Processor, amount float64) {
    for _, processor := range processors {
        processor.ProcessPayment(amount)  // Works with any adapted system
    }
}
```

## 🎨 Advanced Features Demonstrated

### 1. **Factory Pattern Integration**
```go
factory := &PaymentAdapterFactory{}
processor, err := factory.CreateAdapter("paypal")
if err == nil {
    processor.ProcessPayment(100.00)
}
```

### 2. **Composition Adapters**
```go
// UniversalMediaAdapter can handle multiple formats
universalPlayer := NewUniversalMediaAdapter()
universalPlayer.Play("song.mp3")  // Automatically selects MP3 adapter
universalPlayer.Play("song.wav")  // Automatically selects WAV adapter
```

### 3. **Error Handling & Validation**
```go
func (p *PayPalAdapter) ProcessPayment(amount float64) error {
    if amount <= 0 {
        return &payment.PaymentError{
            Code:    "INVALID_AMOUNT",
            Message: "Amount must be greater than zero",
        }
    }
    return p.gateway.MakePayment(amount, "USD")
}
```

## 🔍 When to Use the Adapter Pattern

### ✅ Use When:
- **Legacy Integration**: You need to use existing classes with incompatible interfaces
- **Third-party Libraries**: External libraries don't match your application's interface
- **Interface Standardization**: You want to create a uniform interface for similar functionality
- **Gradual Migration**: Moving from old systems to new ones incrementally

### ❌ Avoid When:
- **Simple Wrappers**: You're just wrapping methods without interface incompatibility
- **Over-engineering**: The interfaces are already compatible
- **Performance Critical**: The extra layer adds unnecessary overhead

## 🏛️ Architectural Benefits

### 1. **Separation of Concerns**
- **Legacy systems** remain unchanged and focused on their original purpose
- **Adapters** handle only the interface translation
- **Modern interfaces** define clean contracts

### 2. **Testability**
```go
// Easy to mock adapters for testing
type MockPaymentAdapter struct{}
func (m *MockPaymentAdapter) ProcessPayment(amount float64) error {
    return nil // Simulate successful payment
}

func TestPaymentProcessing(t *testing.T) {
    mockProcessor := &MockPaymentAdapter{}
    err := mockProcessor.ProcessPayment(100.00)
    assert.NoError(t, err)
}
```

### 3. **Extensibility**
Adding new payment systems is straightforward:
```go
// Add new legacy system
type ApplePayGateway struct { /* implementation */ }

// Create corresponding adapter
type ApplePayAdapter struct {
    gateway *ApplePayGateway
}

// Implement the standard interface
func (a *ApplePayAdapter) ProcessPayment(amount float64) error {
    return a.gateway.ProcessApplePayment(amount)
}
```

## 📚 Related Patterns

- **Bridge Pattern**: Similar structure but different intent (abstraction vs adaptation)
- **Decorator Pattern**: Adds behavior; Adapter changes interface
- **Facade Pattern**: Simplifies interface; Adapter makes incompatible interfaces compatible
- **Proxy Pattern**: Controls access; Adapter enables access

## 🧪 Testing the Implementation

```bash
# Run tests (if you add test files)
go test ./...

# Run with verbose output
go test -v ./...

# Test specific package
go test ./internal/adapters

# Run benchmarks
go test -bench=. ./...
```

## 🔧 Extending the Example

### Adding New Payment Systems
1. Create the legacy system in `internal/legacy/`
2. Implement the adapter in `internal/adapters/`
3. Update the factory to support the new type
4. Add demonstration code in `main.go`

### Adding New Interfaces
1. Define the target interface in `internal/modern/`
2. Create legacy implementations in `internal/legacy/`
3. Build adapters in `internal/adapters/`
4. Demonstrate usage in `main.go`

## 📖 Learning Resources

- [Go Design Patterns](https://github.com/tmrts/go-patterns)
- [Gang of Four Design Patterns](https://en.wikipedia.org/wiki/Design_Patterns)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Interfaces](https://tour.golang.org/methods/9)

---

This implementation demonstrates how the Adapter pattern enables clean integration of legacy systems while maintaining modern, consistent interfaces throughout your Go application. The pattern promotes code reusability, maintainability, and testability while allowing gradual system modernization.