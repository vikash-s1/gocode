# Command Design Pattern in Go

## What is the Command Design Pattern?

The Command pattern is a behavioral design pattern that turns a request into a stand-alone object containing all information about the request. This transformation allows you to parameterize methods with different requests, delay or queue a request's execution, and support undoable operations.

The pattern encapsulates a request as an object, thereby letting you:
- Parameterize clients with different requests
- Queue or log requests
- Support undo operations
- Decouple the object that invokes the operation from the object that performs it

## When to Use Command Pattern

Use the Command pattern when:

1. **Undo/Redo Operations**: When you need to support undo and redo functionality
2. **Queuing Operations**: When you want to queue operations, schedule their execution, or execute them remotely
3. **Logging and Auditing**: When you need to log operations for auditing or crash recovery
4. **Macro Recording**: When you want to record a sequence of operations and replay them
5. **Decoupling**: When you want to decouple the object that invokes the operation from the object that performs it
6. **Parameterizing Objects**: When you want to parameterize objects with operations

### Common Use Cases:
- GUI buttons and menu items
- Remote controls and smart home systems
- Database transactions and rollbacks
- Text editors with undo/redo
- Game action systems
- API request handling
- Batch processing systems

## Implementation Details

The Command pattern typically involves:

1. **Command Interface**: Declares a method for executing a command
2. **Concrete Command**: Implements the command interface and defines the binding between a Receiver and an action
3. **Receiver**: Knows how to perform the operations associated with carrying out a request
4. **Invoker**: Asks the command to carry out the request
5. **Client**: Creates concrete command objects and sets their receivers

### Key Benefits:
- **Decoupling**: Separates the object that invokes the operation from the object that performs it
- **Flexibility**: Easy to add new commands without changing existing code
- **Undo/Redo**: Natural support for undoable operations
- **Logging**: Commands can be logged for auditing or recovery
- **Queuing**: Commands can be queued and executed later
- **Macro Support**: Multiple commands can be combined into composite commands

### Implementation Approaches:
1. **Simple Command**: Basic command with execute method
2. **Undoable Command**: Command with both execute and undo methods
3. **Composite Command**: Command that contains multiple sub-commands
4. **Queued Command**: Commands stored in a queue for batch execution

## Demo

This implementation demonstrates a smart home automation system with various devices:

```bash
cd gocode/CommandDesignPattern
go run main.go
```

The demo shows:
- Basic device control (lights, fan, stereo)
- Remote control with multiple buttons
- Undo/redo functionality
- Macro commands (multiple operations)
- Command queuing and batch execution

## Key Features

- **Device Control**: Control various smart home devices
- **Remote Control**: Simulate a universal remote with programmable buttons
- **Undo Support**: All commands support undo operations
- **Macro Commands**: Combine multiple commands into one
- **Command History**: Track and replay command history
- **Flexible Architecture**: Easy to add new devices and commands

## Summary

The Command pattern is essential for creating flexible, undoable, and loggable operations. It's particularly useful in Go for building systems that need to decouple request senders from receivers, support undo operations, or queue operations for later execution.