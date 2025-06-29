# Strategy Design Pattern in Go

## Overview

The Strategy Design Pattern is a behavioral design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable. The strategy lets the algorithm vary independently from clients that use it.

## Problem It Solves

Without the Strategy pattern, you'd typically handle different behaviors using large conditional statements:

```go
// BAD: Complex conditional logic
func processPayment(paymentType string, amount float64) error {
    if paymentType == "creditcard" {
        // Credit card logic
        validateCard()
        chargeCard(amount)
    } else if paymentType == "paypal" {
        // PayPal logic
        authenticatePayPal()
        transferMoney(amount)
    } else if paymentType == "crypto" {
        // Crypto logic
        connectBlockchain()
        sendTransaction(amount)
    }
    // ... more conditions
}
```

This leads to:
- **Violation of Open/Closed Principle**: Adding new payment methods requires modifying existing code
- **Complex maintenance**: All payment logic mixed together
- **Difficult testing**: Hard to test individual payment methods in isolation
- **Runtime inflexibility**: Cannot change payment methods dynamically

## Solution

The Strategy pattern encapsulates each algorithm (payment method) in separate strategy classes and makes them interchangeable at runtime.

## Implementation Structure

### Core Components

1. **Strategy Interface** (`PaymentStrategy`)
   - Defines the contract for all concrete strategies
   - Ensures consistent behavior across different implementations

2. **Context** (`PaymentContext`) 
   - Maintains a reference to the current strategy
   - Delegates algorithm execution to the strategy object
   - Provides interface for clients to interact with strategies

3. **Concrete Strategies**
   - `CreditCardStrategy`: Handles credit card payments
   - `PayPalStrategy`: Handles PayPal payments  
   - `CryptoStrategy`: Handles cryptocurrency payments
   - `BankTransferStrategy`: Handles bank transfer payments

## Payment System Example

Our implementation simulates a payment processing system with multiple payment methods:

```
┌─────────────────┐
│ PaymentContext  │
│                 │     ┌─────────────────────┐
│ - strategy      │────▶│ PaymentStrategy     │
│ - customer      │     │                     │
│                 │     │ + Pay()             │
│ + SetStrategy() │     │ + ValidatePayment() │
│ + ProcessPayment│     │ + GetPaymentMethod()│
└─────────────────┘     └─────────────────────┘
                                   ▲
                    ┌──────────────┼──────────────┐
                    │              │              │
        ┌───────────────────┐ ┌──────────────┐ ┌─────────────────┐
        │CreditCardStrategy │ │PayPalStrategy│ │  CryptoStrategy │
        │                  │ │              │ │                 │
        │+ Pay()           │ │+ Pay()       │ │+ Pay()          │
        │+ ValidatePayment()│ │+ ValidatePayment()│+ ValidatePayment()│
        └───────────────────┘ └──────────────┘ └─────────────────┘
```

## Step-by-Step Implementation

### Step 1: Define the Strategy Interface

```go
type PaymentStrategy interface {
    Pay(amount float64) error
    GetPaymentMethod() string
    ValidatePayment(amount float64) error
}
```

**Purpose**: Establishes the contract that all payment strategies must implement, ensuring consistency and interchangeability.

### Step 2: Create the Context

```go
type PaymentContext struct {
    strategy PaymentStrategy
    customer string
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
    pc.strategy = strategy
}

func (pc *PaymentContext) ProcessPayment(amount float64) error {
    return pc.strategy.Pay(amount)
}
```

**Key Features**:
- Holds reference to current strategy
- Allows runtime strategy switching
- Delegates payment processing to the strategy

### Step 3: Implement Concrete Strategies

#### CreditCardStrategy
```go
type CreditCardStrategy struct {
    CardNumber string
    CVV        string
    ExpiryDate string
    CardHolder string
}

func (cc *CreditCardStrategy) Pay(amount float64) error {
    // Credit card specific payment logic
    fmt.Printf("Processing credit card payment: $%.2f\n", amount)
    // Simulate payment processing...
    return nil
}
```

**Characteristics**:
- Encapsulates credit card payment logic
- Validates card details
- Handles card-specific errors

#### PayPalStrategy
```go
type PayPalStrategy struct {
    Email    string
    Password string
}

func (pp *PayPalStrategy) Pay(amount float64) error {
    // PayPal specific payment logic
    fmt.Printf("Processing PayPal payment: $%.2f\n", amount)
    // Simulate PayPal authentication and payment...
    return nil
}
```

**Characteristics**:
- Handles PayPal authentication
- Manages PayPal-specific workflows
- Different validation rules than credit cards

## Key Benefits Demonstrated

### 1. **Runtime Strategy Switching**
```go
paymentContext := NewPaymentContext("John Doe")

// Start with credit card
paymentContext.SetStrategy(NewCreditCardStrategy(...))
paymentContext.ProcessPayment(100.00)

// Switch to PayPal at runtime
paymentContext.SetStrategy(NewPayPalStrategy(...))
paymentContext.ProcessPayment(50.00)
```

### 2. **Open/Closed Principle**
Adding new payment methods doesn't require modifying existing code:
```go
// Easy to add new strategies
type ApplePayStrategy struct {
    TouchID string
    DeviceID string
}

func (ap *ApplePayStrategy) Pay(amount float64) error {
    // Apple Pay specific logic
    return nil
}
```

### 3. **Single Responsibility Principle**
Each strategy handles only its specific payment method:
```go
// CreditCardStrategy only handles credit card logic
func (cc *CreditCardStrategy) ValidatePayment(amount float64) error {
    if len(cc.CardNumber) != 16 {
        return fmt.Errorf("invalid card number")
    }
    // Credit card specific validation...
}
```

### 4. **Elimination of Conditional Logic**
Instead of complex if/else chains, we have clean delegation:
```go
// GOOD: Clean delegation
func (pc *PaymentContext) ProcessPayment(amount float64) error {
    return pc.strategy.Pay(amount)
}
```

## Running the Example

### Build and Run
```bash
cd StrategyDesignPattern
go run main.go
```

### Interactive Mode Features
1. **Customer Setup**: Enter customer name
2. **Strategy Selection**: Choose from 4 payment methods
3. **Runtime Switching**: Change payment methods dynamically
4. **Payment Processing**: Process payments with current strategy
5. **Error Handling**: See validation and processing errors

### Automated Demos
The program demonstrates:
1. **All Payment Strategies**: Shows each payment method in action
2. **Runtime Switching**: Demonstrates changing strategies during execution
3. **Error Scenarios**: Simulates payment failures and validation errors
4. **Different Amounts**: Processes various payment amounts

## Strategy Comparison

| Strategy | Processing Time | Failure Rate | Special Features |
|----------|----------------|--------------|------------------|
| Credit Card | 1 second | 10% | Card validation, CVV check |
| PayPal | 0.8 seconds | 5% | Email authentication |
| Cryptocurrency | 1.5 seconds | 15% | Blockchain processing |
| Bank Transfer | 2 seconds | 8% | Account verification |

## When to Use Strategy Pattern

✅ **Use When**:
- You have multiple ways to perform a task
- You want to switch algorithms at runtime
- You have complex conditional logic based on types
- You need to add new algorithms frequently
- Algorithms are independent and interchangeable

❌ **Avoid When**:
- You only have one or two simple algorithms
- Algorithms are tightly coupled to context
- Performance overhead of indirection is critical
- Algorithms rarely change

## Real-World Applications

### E-commerce Systems
```go
// Shipping strategies
type ShippingStrategy interface {
    CalculateCost(weight float64, distance float64) float64
    GetDeliveryTime() time.Duration
}

// StandardShipping, ExpressShipping, OvernightShipping
```

### Game Development
```go
// AI behavior strategies
type AIStrategy interface {
    MakeMove(gameState *GameState) Move
}

// AggressiveAI, DefensiveAI, RandomAI
```

### Data Processing
```go
// Compression strategies
type CompressionStrategy interface {
    Compress(data []byte) []byte
    Decompress(data []byte) []byte
}

// GzipCompression, ZipCompression, LZ4Compression
```

### Sorting Algorithms
```go
// Sorting strategies
type SortStrategy interface {
    Sort(data []int) []int
}

// QuickSort, MergeSort, BubbleSort
```

## Comparison with Other Patterns

### vs State Pattern
- **Strategy**: Client chooses algorithm, strategies are independent
- **State**: Object behavior changes based on internal state, states may know about each other

### vs Command Pattern
- **Strategy**: Encapsulates algorithms/behaviors
- **Command**: Encapsulates requests/actions

### vs Template Method Pattern
- **Strategy**: Uses composition, entire algorithm is replaceable
- **Template Method**: Uses inheritance, only parts of algorithm vary

## Advanced Considerations

### Strategy Factory
For complex strategy creation:
```go
type PaymentStrategyFactory struct{}

func (f *PaymentStrategyFactory) CreateStrategy(strategyType string, config map[string]string) PaymentStrategy {
    switch strategyType {
    case "creditcard":
        return NewCreditCardStrategy(config["cardNumber"], config["cvv"], ...)
    case "paypal":
        return NewPayPalStrategy(config["email"], config["password"])
    // ... other strategies
    }
}
```

### Strategy Configuration
For configurable strategies:
```go
type StrategyConfig struct {
    RetryAttempts int
    Timeout       time.Duration
    EnableLogging bool
}

type ConfigurableStrategy struct {
    config StrategyConfig
    // ... other fields
}
```

### Strategy Chaining
For composite strategies:
```go
type ChainedPaymentStrategy struct {
    primary   PaymentStrategy
    fallback  PaymentStrategy
}

func (cps *ChainedPaymentStrategy) Pay(amount float64) error {
    if err := cps.primary.Pay(amount); err != nil {
        return cps.fallback.Pay(amount)
    }
    return nil
}
```

## Performance Considerations

### Memory Usage
- Each strategy object consumes memory
- Consider object pooling for frequently used strategies

### Execution Overhead
- Minimal overhead from interface calls
- Strategy switching is O(1) operation

### Optimization Tips
```go
// Pre-create strategies to avoid allocation overhead
var (
    creditCardStrategy = NewCreditCardStrategy(...)
    paypalStrategy     = NewPayPalStrategy(...)
)

func GetStrategy(strategyType string) PaymentStrategy {
    switch strategyType {
    case "creditcard":
        return creditCardStrategy
    case "paypal":
        return paypalStrategy
    }
}
```

This implementation demonstrates how the Strategy pattern creates flexible, maintainable code by encapsulating algorithms and making them interchangeable at runtime, eliminating complex conditional logic and adhering to SOLID principles.