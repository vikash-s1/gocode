# Go Programming Guide

Go (also known as Golang) is an open-source programming language developed by Google. It's designed for simplicity, efficiency, and excellent concurrency support.

## Getting Started

### Installation

1. Download Go from [golang.org](https://golang.org/dl/)
2. Follow the installation instructions for your operating system
3. Verify installation:
   ```bash
   go version
   ```

### Your First Go Program

Create a file named `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run it:
```bash
go run main.go
```

## Key Features

- **Simple syntax** - Easy to learn and read
- **Fast compilation** - Compiles to native machine code
- **Built-in concurrency** - Goroutines and channels
- **Garbage collection** - Automatic memory management
- **Static typing** - Type safety with inference
- **Cross-platform** - Compile for multiple architectures

## Basic Syntax

### Variables
```go
var name string = "Go"
age := 25  // Short declaration
const pi = 3.14159
```

### Functions
```go
func add(a, b int) int {
    return a + b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### Control Structures
```go
// If statement
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x == 10 {
    fmt.Println("x equals 10")
} else {
    fmt.Println("x is less than 10")
}

// For loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Range loop
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Data Types

### Basic Types
- `bool`
- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `complex64`, `complex128`

### Composite Types
```go
// Array
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice
slice := []int{1, 2, 3, 4, 5}
slice = append(slice, 6)

// Map
m := make(map[string]int)
m["apple"] = 5
m["banana"] = 3

// Struct
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```

## Concurrency

### Goroutines
```go
func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello()  // Run in goroutine
    time.Sleep(1 * time.Second)  // Wait for goroutine
}
```

### Channels
```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from channel!"
    }()
    
    message := <-ch
    fmt.Println(message)
}
```

## Package Management

### Go Modules
Initialize a new module:
```bash
go mod init [module-name]
```

Add dependencies:
```bash
go get [package-name]
```

Common commands:
```bash
go build      # Compile the program
go run        # Compile and run
go test       # Run tests
go fmt        # Format code
go vet        # Examine code for errors
go mod tidy   # Clean up dependencies
```

## Project Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── internal/
│   └── handlers/
├── pkg/
│   └── utils/
└── cmd/
    └── server/
```

## Testing

Create test files with `_test.go` suffix:

```go
// math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
```

Run tests:
```bash
go test
go test -v  # Verbose output
```

## Best Practices

1. **Follow Go conventions** - Use `gofmt` and `golint`
2. **Handle errors explicitly** - Don't ignore error returns
3. **Use meaningful names** - Clear, descriptive variable and function names
4. **Keep functions small** - Single responsibility principle
5. **Use interfaces** - Program to interfaces, not implementations
6. **Document your code** - Use comments for exported functions

## Useful Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Playground](https://play.golang.org/)
- [Awesome Go](https://github.com/avelino/awesome-go)

## Common Libraries

- **Web frameworks**: Gin, Echo, Fiber
- **Database**: GORM, sqlx
- **Testing**: Testify, GoMock
- **CLI**: Cobra, urfave/cli
- **HTTP client**: Resty
- **Logging**: Logrus, Zap

Happy coding with Go! 🚀