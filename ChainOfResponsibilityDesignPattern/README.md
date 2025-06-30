# Chain of Responsibility Design Pattern in Go

## What is the Chain of Responsibility Pattern?

The Chain of Responsibility is a behavioral design pattern that allows you to pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

### Key Characteristics:
- **Decoupling**: The sender of a request is decoupled from its receivers
- **Dynamic Chain**: Handlers can be added, removed, or reordered at runtime
- **Single Responsibility**: Each handler has a specific responsibility
- **Flexible Processing**: Multiple handlers can process the same request

## When to Use Chain of Responsibility

### Use this pattern when:

1. **Multiple Objects Can Handle a Request**
   - When you have several objects that can process a request, but you don't know which one will handle it until runtime

2. **Dynamic Handler Selection**
   - When you want to specify handlers dynamically or change the order of processing

3. **Avoiding Tight Coupling**
   - When you want to decouple request senders from receivers

4. **Processing Pipeline**
   - When you need to process requests through a series of filters or validators

5. **Event Handling Systems**
   - GUI event handling, middleware processing, or logging systems

### Common Use Cases:
- **Logging Systems**: Different log levels (INFO, WARNING, ERROR)
- **Authentication/Authorization**: Multiple authentication methods
- **Request Processing**: Web middleware, API gateways
- **Validation Chains**: Form validation, data processing
- **Event Handling**: GUI events, game event systems

## Implementation Details

### Core Components:

1. **Handler Interface**: Defines the contract for all handlers
2. **Base Handler**: Provides common functionality and chain management
3. **Concrete Handlers**: Implement specific processing logic
4. **Request**: Contains the data to be processed

### Key Implementation Points:

- **Handler Interface**: All handlers implement the same interface
- **Chain Management**: Each handler holds a reference to the next handler
- **Processing Logic**: Handlers decide whether to process or pass the request
- **Flexible Chain**: The chain can be built dynamically at runtime

## Project Structure

```
ChainOfResponsibilityDesignPattern/
├── go.mod
├── main.go
├── internal/
│   └── handlers/
│       ├── handler.go          # Handler interface and Request struct
│       ├── base_handler.go     # Base handler with common functionality
│       ├── info_handler.go     # Handles INFO level requests
│       ├── warning_handler.go  # Handles WARNING level requests
│       └── error_handler.go    # Handles ERROR level requests
└── README.md
```

## How It Works

1. **Chain Setup**: Handlers are linked together in a specific order
2. **Request Processing**: A request is sent to the first handler in the chain
3. **Handler Decision**: Each handler decides whether to:
   - Process the request and stop the chain
   - Process the request and continue the chain
   - Skip processing and pass to the next handler
4. **Chain Traversal**: The request moves through the chain until handled or the chain ends

## Running the Demo

```bash
cd gocode/ChainOfResponsibilityDesignPattern
go run main.go
```

### Expected Output:
The demo shows how different types of requests (INFO, WARNING, ERROR) are processed by appropriate handlers in the chain.

## Advantages

- **Flexibility**: Easy to add, remove, or reorder handlers
- **Decoupling**: Sender doesn't need to know which handler will process the request
- **Single Responsibility**: Each handler has one specific purpose
- **Runtime Configuration**: Chain can be modified at runtime

## Disadvantages

- **Performance**: Request might traverse the entire chain
- **Debugging**: Can be harder to debug the flow through multiple handlers
- **No Guarantee**: No guarantee that a request will be handled

## Real-World Examples

- **Web Middleware**: Express.js middleware, HTTP interceptors
- **Logging Frameworks**: Log4j, Logrus with different appenders
- **Authentication Systems**: OAuth, JWT, Basic Auth fallbacks
- **Validation Pipelines**: Form validation, data sanitization
- **Event Systems**: DOM event bubbling, game event handling

## Summary

The Chain of Responsibility pattern provides a clean way to handle requests through a series of processors without tightly coupling the sender to specific receivers. It's particularly useful in scenarios where you need flexible, configurable processing pipelines that can be modified at runtime.