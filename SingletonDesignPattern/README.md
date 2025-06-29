# Singleton Design Pattern in Go

## Overview

The Singleton pattern ensures that a class has only one instance and provides a global point of access to that instance. This pattern is useful when you need exactly one instance of a class to coordinate actions across the system.

## When to Use Singleton

- **Database connections**: When you want to share a single database connection pool
- **Logging**: When you need a centralized logging mechanism
- **Configuration management**: When you need global access to application settings
- **Cache management**: When you need a single cache instance across the application

## Implementation Details

### Key Components

1. **Private constructor**: Prevents direct instantiation
2. **Static instance variable**: Holds the single instance
3. **Static method**: Provides global access to the instance
4. **Thread safety**: Ensures safe creation in concurrent environments

### Thread-Safe Implementation

In Go, we use `sync.Once` to ensure thread-safe singleton creation:

```go
var (
    instance *DatabaseConnection
    once sync.Once
)

func GetInstance() *DatabaseConnection {
    once.Do(func() {
        instance = &DatabaseConnection{
            connectionString: "localhost:5432/mydb",
            isConnected:      false,
        }
    })
    return instance
}
```

## Project Structure

```
SingletonDesignPattern/
├── go.mod
├── main.go
├── README.md
└── internal/
    └── singleton/
        ├── database.go    # Database connection singleton
        └── logger.go      # Logger singleton
```

## Code Walkthrough

### Step 1: Database Singleton (`internal/singleton/database.go`)

- **Private struct**: `DatabaseConnection` represents our singleton
- **Package-level variables**: `instance` and `once` for thread-safe creation
- **GetInstance()**: Returns the singleton instance using `sync.Once`
- **Business methods**: `Connect()`, `Disconnect()`, etc.

Key features:
- Thread-safe initialization using `sync.Once`
- Private instance variable prevents external access
- Single point of access through `GetInstance()`

### Step 2: Logger Singleton (`internal/singleton/logger.go`)

- **Concurrent-safe logging**: Uses `sync.Mutex` for thread-safe operations
- **Separate sync.Once**: Each singleton has its own creation mechanism
- **Log level management**: Demonstrates state management in singleton

Key features:
- Thread-safe logging operations
- Configurable log levels
- Timestamp formatting

### Step 3: Main Application (`main.go`)

The main function demonstrates three key aspects:

#### 1. Basic Singleton Usage
```go
db1 := singleton.GetInstance()
db2 := singleton.GetInstance()
fmt.Printf("Same instance? %t\n", db1 == db2) // true
```

#### 2. State Sharing
```go
db1.Connect()
fmt.Printf("DB2 connected: %t\n", db2.IsConnected()) // true
```

#### 3. Thread Safety
```go
// Multiple goroutines getting the same instance
for i := 0; i < 10; i++ {
    go func(index int) {
        instances[index] = singleton.GetInstance()
    }(i)
}
```

## Running the Code

```bash
cd gocode/SingletonDesignPattern
go run main.go
```

## Expected Output

```
=== Singleton Design Pattern Demo ===

1. Database Singleton Demo:
Creating new database connection instance
DB1 Connection String: localhost:5432/mydb
DB2 Connection String: localhost:5432/mydb
Are db1 and db2 the same instance? true
Connected to database: localhost:5432/mydb
DB2 is connected: true
Disconnected from database
DB1 is connected: false

--------------------------------------------------

2. Logger Singleton Demo:
Creating new logger instance
[2024-01-15 10:30:45] INFO: First log message
[2024-01-15 10:30:45] DEBUG: Second log message
Are logger1 and logger2 the same instance? true
Log level set to: ERROR
Logger2 log level: ERROR

--------------------------------------------------

3. Thread Safety Demo:
Goroutine 0 got instance
Goroutine 1 got instance
...
All 10 instances are the same: true
```

## Advantages

1. **Controlled access**: Single point of access to the instance
2. **Memory efficiency**: Only one instance exists in memory
3. **Global state**: Shared state across the application
4. **Lazy initialization**: Instance created only when needed

## Disadvantages

1. **Global state**: Can make testing difficult
2. **Hidden dependencies**: Makes dependencies less explicit
3. **Scalability**: Can become a bottleneck in highly concurrent applications
4. **Violation of SRP**: Often handles both creation and business logic

## Best Practices

1. **Use sparingly**: Only when you truly need a single instance
2. **Thread safety**: Always ensure thread-safe implementation
3. **Interface-based**: Consider using interfaces for better testability
4. **Avoid global state**: Consider dependency injection as an alternative

## Alternative Approaches

Instead of Singleton, consider:
- **Dependency Injection**: Pass instances through constructors
- **Factory Pattern**: Control instance creation without enforcing singularity
- **Service Locator**: Centralized registry for services

## Testing Considerations

Singletons can make unit testing challenging because:
- Global state persists between tests
- Hard to mock or replace for testing
- Tests may have hidden dependencies

Consider using interfaces and dependency injection for better testability.