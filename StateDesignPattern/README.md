# State Design Pattern in Go

## Overview

The State Design Pattern is a behavioral design pattern that allows an object to change its behavior when its internal state changes. The object appears to change its class, making it seem like the object's behavior is tied to its state.

## Problem It Solves

Without the State pattern, you'd typically use large conditional statements (if/else or switch) to handle different behaviors based on the object's state. This leads to:
- Complex, hard-to-maintain code
- Violation of Open/Closed Principle
- Difficulty adding new states
- State-specific logic scattered throughout the codebase

## Solution

The State pattern encapsulates state-specific behavior into separate state classes and delegates behavior to the current state object.

## Implementation Structure

### Core Components

1. **State Interface** (`State`)
   - Defines the contract for all concrete states
   - Each method represents an action that can be performed

2. **Context** (`VendingMachine`)
   - Maintains a reference to the current state
   - Delegates state-specific behavior to the current state object
   - Provides methods for state transitions

3. **Concrete States**
   - `IdleState`: Initial state, waiting for coin
   - `CoinInsertedState`: Coin inserted, waiting for product selection
   - `ProductSelectedState`: Product selected, dispensing in progress
   - `OutOfStockState`: No products available

## Vending Machine Example

Our implementation simulates a vending machine with the following states and transitions:

```
┌─────────────┐    InsertCoin()    ┌──────────────────┐
│ IdleState   │ ──────────────────▶│ CoinInsertedState│
│             │                    │                  │
└─────────────┘                    └──────────────────┘
       ▲                                     │
       │                                     │ SelectProduct()
       │ Cancel()                            ▼
       │                           ┌──────────────────┐
       │                           │ProductSelectedState│
       │                           │                  │
       │                           └──────────────────┘
       │                                     │
       │ DispenseProduct()                   │ DispenseProduct()
       │ (if products available)             │ (if no products left)
       │                                     ▼
       │                           ┌──────────────────┐
       └───────────────────────────│ OutOfStockState  │
                                   │                  │
                                   └──────────────────┘
```

## Step-by-Step Implementation

### Step 1: Define the State Interface

```go
type State interface {
    InsertCoin(machine *VendingMachine)
    SelectProduct(machine *VendingMachine)
    DispenseProduct(machine *VendingMachine)
    Cancel(machine *VendingMachine)
    GetStateName() string
}
```

**Purpose**: Establishes the contract that all concrete states must implement.

### Step 2: Create the Context (VendingMachine)

```go
type VendingMachine struct {
    currentState State
    coinInserted bool
    productCount int
}
```

**Key Methods**:
- `SetState()`: Changes the current state
- `InsertCoin()`, `SelectProduct()`, etc.: Delegate to current state
- Helper methods for state management

### Step 3: Implement Concrete States

#### IdleState
- **Behavior**: Accepts coins, rejects other operations
- **Transitions**: To `CoinInsertedState` when coin is inserted

#### CoinInsertedState  
- **Behavior**: Accepts product selection or cancellation
- **Transitions**: 
  - To `ProductSelectedState` when product is selected
  - To `IdleState` when cancelled

#### ProductSelectedState
- **Behavior**: Handles product dispensing
- **Transitions**:
  - To `IdleState` if products remain
  - To `OutOfStockState` if no products left

#### OutOfStockState
- **Behavior**: Rejects all operations
- **Transitions**: None (terminal state in this example)

## Key Benefits Demonstrated

### 1. **Single Responsibility Principle**
Each state class handles only its specific behavior:
```go
func (s *IdleState) InsertCoin(machine *VendingMachine) {
    println("💰 Coin inserted! Please select a product.")
    machine.SetCoinInserted(true)
    machine.SetState(&CoinInsertedState{})
}
```

### 2. **Open/Closed Principle**
Adding new states doesn't require modifying existing code:
```go
// Easy to add new states like MaintenanceState
type MaintenanceState struct{}
```

### 3. **Eliminates Complex Conditionals**
Instead of:
```go
// BAD: Complex conditional logic
func (vm *VendingMachine) InsertCoin() {
    if vm.state == "idle" {
        // handle idle logic
    } else if vm.state == "coinInserted" {
        // handle coin inserted logic
    } else if vm.state == "productSelected" {
        // handle product selected logic
    }
    // ... more conditions
}
```

We have:
```go
// GOOD: Clean delegation
func (vm *VendingMachine) InsertCoin() {
    vm.currentState.InsertCoin(vm)
}
```

## Running the Example

### Build and Run
```bash
cd StateDesignPattern
go run main.go
```

### Interactive Mode
The program starts with an interactive mode where you can:
1. Insert coins
2. Select products
3. Dispense products
4. Cancel transactions
5. Check machine status

### Automated Demos
After the interactive mode, the program runs automated demos showing:
1. Normal purchase flow
2. Cancellation scenarios
3. Invalid operations handling
4. Out-of-stock behavior

## State Transitions in Action

### Normal Flow Example:
```
State: Idle → Insert Coin → CoinInserted → Select Product → ProductSelected → Dispense → Idle
```

### Cancellation Flow:
```
State: Idle → Insert Coin → CoinInserted → Cancel → Idle
```

### Out of Stock Flow:
```
State: ProductSelected → Dispense (last product) → OutOfStock
```

## When to Use State Pattern

✅ **Use When**:
- Object behavior changes significantly based on internal state
- You have complex conditional logic based on state
- State transitions are well-defined
- You need to add new states frequently

❌ **Avoid When**:
- Simple state logic (few states, simple transitions)
- States don't have significantly different behaviors
- Performance is critical (adds indirection overhead)

## Real-World Applications

- **UI Components**: Button states (enabled, disabled, pressed)
- **Game Development**: Character states (idle, running, jumping, attacking)
- **Network Protocols**: Connection states (connecting, connected, disconnected)
- **Document Workflow**: Draft, review, approved, published states
- **Order Processing**: Pending, confirmed, shipped, delivered states

## Comparison with Other Patterns

### vs Strategy Pattern
- **State**: Behavior changes based on internal state, states know about each other
- **Strategy**: Behavior chosen by client, strategies are independent

### vs Command Pattern  
- **State**: Encapsulates behavior based on object state
- **Command**: Encapsulates requests as objects

## Advanced Considerations

### State Persistence
For real applications, consider:
```go
type VendingMachine struct {
    currentState State
    stateHistory []string  // For audit trails
    // ... other fields
}
```

### Concurrent Access
For multi-threaded environments:
```go
import "sync"

type VendingMachine struct {
    mu           sync.RWMutex
    currentState State
    // ... other fields
}
```

### State Factory
For complex state creation:
```go
type StateFactory struct{}

func (sf *StateFactory) CreateState(stateType string) State {
    switch stateType {
    case "idle":
        return &IdleState{}
    case "coinInserted":
        return &CoinInsertedState{}
    // ... other states
    }
}
```

This implementation demonstrates how the State pattern creates clean, maintainable code by encapsulating state-specific behavior and eliminating complex conditional logic.