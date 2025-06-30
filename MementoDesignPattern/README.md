# Memento Design Pattern in Go

## What is the Memento Design Pattern?

The Memento pattern is a behavioral design pattern that allows you to save and restore the previous state of an object without revealing the details of its implementation. It provides the ability to restore an object to its previous state (undo via rollback).

The pattern captures and externalizes an object's internal state so that the object can be restored to this state later, all without violating encapsulation. It's essentially a snapshot of an object's state at a particular point in time.

## When to Use Memento Pattern

Use the Memento pattern when:

1. **Undo Operations**: When you need to implement undo functionality in your application
2. **Checkpoints**: When you want to save snapshots of an object's state at certain points
3. **State Restoration**: When you need to restore an object to a previous state after an operation fails
4. **History Tracking**: When you need to maintain a history of state changes
5. **Rollback Functionality**: When you need to rollback to a previous stable state
6. **Encapsulation Preservation**: When you want to save state without breaking encapsulation

### Common Use Cases:
- Text editors with undo/redo functionality
- Game save states and checkpoints
- Database transaction rollbacks
- Configuration management systems
- Version control systems
- Drawing applications with undo
- Form data backup and restoration
- State machines with rollback capability

## Implementation Details

The Memento pattern typically involves three key participants:

1. **Originator**: The object whose state needs to be saved and restored
2. **Memento**: A snapshot of the originator's state at a specific point in time
3. **Caretaker**: Manages mementos but never operates on or examines their contents

### Key Benefits:
- **Encapsulation**: Preserves encapsulation boundaries by not exposing internal state
- **Simplicity**: Simplifies the originator by delegating state management
- **Undo/Redo**: Natural support for undo and redo operations
- **Snapshots**: Easy creation of object state snapshots
- **Recovery**: Enables recovery from failed operations

### Implementation Approaches:
1. **Simple Memento**: Basic state capture and restoration
2. **Multiple Mementos**: Managing multiple snapshots with history
3. **Incremental Mementos**: Storing only changes between states
4. **Compressed Mementos**: Optimized storage for large states

## Demo

This implementation demonstrates a text editor with undo/redo functionality:

```bash
cd gocode/MementoDesignPattern
go run main.go
```

The demo shows:
- Text editor with content manipulation
- Save and restore functionality
- Undo/redo operations
- History management
- Multiple checkpoint support

## Key Features

- **Text Editor**: Complete text editor with content management
- **State Snapshots**: Save editor state at any point
- **Undo/Redo**: Full undo and redo functionality
- **History Management**: Maintain history of all changes
- **Encapsulation**: Internal state remains private
- **Multiple Formats**: Support for different content types

## Summary

The Memento pattern is essential for implementing undo functionality and state management in applications. It provides a clean way to capture and restore object states while maintaining encapsulation, making it particularly useful in Go applications that need robust state management capabilities.