# Factory Design Pattern in Go

## Overview

The Factory Design Pattern is a creational design pattern that provides an interface for creating objects without specifying their exact classes. Instead of calling constructors directly, the pattern uses a factory method to create objects based on input parameters.

## Why Use Factory Pattern?

- **Encapsulation**: Object creation logic is centralized and hidden from clients
- **Flexibility**: Easy to add new product types without modifying existing code
- **Loose Coupling**: Clients depend on interfaces, not concrete implementations
- **Single Responsibility**: Creation logic is separated from business logic

## Project Structure

```
FactoryDesignPattern/
├── main.go                          # Main application demonstrating the pattern
├── go.mod                          # Go module file
└── internal/
    └── products/
        ├── vehicle.go              # Product interfaces and concrete implementations
        └── factory.go              # Factory implementation
```

## Implementation Details

### 1. Product Interface (`Vehicle`)

The `Vehicle` interface defines the contract that all concrete products must implement:

```go
type Vehicle interface {
    Start() string
    Stop() string
    GetInfo() string
}
```

### 2. Concrete Products

Three concrete implementations of the `Vehicle` interface:

- **Car**: Basic vehicle with brand and model
- **Motorcycle**: Two-wheeled vehicle implementation
- **Truck**: Heavy vehicle with additional capacity property

### 3. Factory (`VehicleFactory`)

The factory encapsulates the object creation logic:

```go
type VehicleFactory struct{}

func (vf *VehicleFactory) CreateVehicle(vehicleType VehicleType, brand, model string) (Vehicle, error)
```

## Step-by-Step Explanation

### Step 1: Define the Product Interface

```go
// Vehicle interface defines common behavior for all vehicles
type Vehicle interface {
    Start() string
    Stop() string
    GetInfo() string
}
```

**Purpose**: Establishes a contract that all vehicle types must follow, ensuring consistency across different implementations.

### Step 2: Implement Concrete Products

```go
type Car struct {
    Brand string
    Model string
}

func (c *Car) Start() string {
    return fmt.Sprintf("%s %s engine started", c.Brand, c.Model)
}
// ... other methods
```

**Purpose**: Each concrete product implements the `Vehicle` interface with its specific behavior while maintaining the same interface.

### Step 3: Create the Factory

```go
type VehicleFactory struct{}

func (vf *VehicleFactory) CreateVehicle(vehicleType VehicleType, brand, model string) (Vehicle, error) {
    switch strings.ToLower(string(vehicleType)) {
    case string(CarType):
        return &Car{Brand: brand, Model: model}, nil
    // ... other cases
    }
}
```

**Purpose**: Centralizes object creation logic and provides a single point of control for instantiating different vehicle types.

### Step 4: Use the Factory

```go
factory := products.NewVehicleFactory()
car, err := factory.CreateVehicle(products.CarType, "Toyota", "Camry")
```

**Purpose**: Clients use the factory to create objects without knowing the specific implementation details.

## Key Components

### VehicleType Enum

```go
type VehicleType string

const (
    CarType        VehicleType = "car"
    MotorcycleType VehicleType = "motorcycle"
    TruckType      VehicleType = "truck"
)
```

**Purpose**: Provides type safety and prevents invalid vehicle type strings.

### Error Handling

The factory returns an error for unsupported vehicle types:

```go
default:
    return nil, errors.New("unknown vehicle type: " + string(vehicleType))
```

**Purpose**: Graceful handling of invalid inputs with clear error messages.

### Extended Factory Methods

```go
func (vf *VehicleFactory) CreateVehicleWithCapacity(vehicleType VehicleType, brand, model string, capacity int) (Vehicle, error)
```

**Purpose**: Demonstrates how factories can be extended with specialized creation methods for specific requirements.

## Running the Example

```bash
cd gocode/FactoryDesignPattern
go run main.go
```

## Expected Output

```
=== Factory Design Pattern Demo ===

1. Creating Different Vehicle Types:
-----------------------------------
✓ Car: Toyota Camry
  Toyota Camry engine started
  Toyota Camry engine stopped

✓ Motorcycle: Harley-Davidson Street 750
  Harley-Davidson Street 750 motorcycle started
  Harley-Davidson Street 750 motorcycle stopped

✓ Truck: Ford F-150 (Capacity: 10 tons)
  Ford F-150 truck engine started
  Ford F-150 truck engine stopped

2. Error Handling:
------------------
✗ Expected error: unknown vehicle type: airplane

3. Supported Vehicle Types:
---------------------------
1. car
2. motorcycle
3. truck

4. Custom Truck with Specific Capacity:
---------------------------------------
✓ Truck: Volvo FH16 (Capacity: 25 tons)
  Volvo FH16 truck engine started
  Volvo FH16 truck engine stopped
```

## Benefits Demonstrated

1. **Extensibility**: Adding new vehicle types requires only implementing the `Vehicle` interface and updating the factory
2. **Maintainability**: Creation logic is centralized in one place
3. **Type Safety**: Using enums prevents runtime errors from invalid types
4. **Error Handling**: Graceful handling of invalid inputs
5. **Interface Segregation**: Clients depend only on the `Vehicle` interface

## Real-World Applications

- **Database Connections**: Creating different database connection types (MySQL, PostgreSQL, MongoDB)
- **Payment Processing**: Creating different payment processors (PayPal, Stripe, Square)
- **File Parsers**: Creating parsers for different file formats (JSON, XML, CSV)
- **UI Components**: Creating different UI elements based on configuration
- **Logging Systems**: Creating different loggers (file, console, remote)

## Variations

- **Abstract Factory**: For creating families of related objects
- **Factory Method**: Using inheritance instead of composition
- **Simple Factory**: A simplified version without the factory interface
- **Registry Pattern**: Combining factory with a registry for dynamic type registration