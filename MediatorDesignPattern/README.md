# Mediator Design Pattern in Go

## What is the Mediator Pattern?

The Mediator is a behavioral design pattern that defines how a set of objects interact with each other. Instead of objects communicating directly, they communicate through a central mediator object. This promotes loose coupling by preventing objects from referring to each other explicitly and letting you vary their interaction independently.

### Key Characteristics:
- **Centralized Communication**: All communication goes through a central mediator
- **Loose Coupling**: Components don't know about each other directly
- **Simplified Relationships**: Reduces many-to-many relationships to one-to-many
- **Reusable Components**: Components can be reused in different contexts

## When to Use Mediator Pattern

### Use this pattern when:

1. **Complex Communication**
   - When you have a set of objects that communicate in complex but well-defined ways

2. **Tight Coupling Issues**
   - When objects are tightly coupled and hard to reuse because they refer to many other objects

3. **Centralized Control**
   - When you want to centralize complex communications and control logic

4. **Reusable Components**
   - When you want to create reusable components that can work in different contexts

5. **Behavior Distribution**
   - When behavior is distributed between several objects and you want to customize it without subclassing

### Common Use Cases:
- **Chat Applications**: Users communicate through a chat room mediator
- **Air Traffic Control**: Aircraft communicate through control tower
- **GUI Components**: Dialog boxes coordinating between form elements
- **Workflow Systems**: Steps in a process coordinated by a workflow engine
- **Event Systems**: Components communicating through event buses

## Implementation Details

### Core Components:

1. **Mediator Interface**: Defines the contract for communication
2. **Concrete Mediator**: Implements the mediation logic and maintains references to components
3. **Component Interface**: Defines the contract for components that use the mediator
4. **Concrete Components**: Implement specific functionality and communicate through mediator

### Key Implementation Points:

- **Interface Segregation**: Clear interfaces for both mediator and components
- **Event-Driven**: Communication happens through events/notifications
- **State Management**: Mediator can maintain shared state
- **Component Registration**: Components register/unregister with mediator

## Project Structure

```
MediatorDesignPattern/
├── go.mod
├── main.go
├── internal/
│   ├── mediator/
│   │   ├── mediator.go           # Mediator and Component interfaces
│   │   └── chat_mediator.go      # Concrete mediator implementation
│   └── components/
│       ├── base_user.go          # Base user functionality
│       ├── regular_user.go       # Regular chat user
│       ├── moderator_user.go     # Moderator with special privileges
│       └── bot_user.go           # Automated bot user
└── README.md
```

## How It Works

1. **Component Registration**: Components register themselves with the mediator
2. **Event Notification**: When a component wants to communicate, it notifies the mediator
3. **Mediation Logic**: Mediator processes the event and determines which components to notify
4. **Event Distribution**: Mediator sends appropriate notifications to relevant components
5. **Component Response**: Components handle notifications and may trigger new events

## Chat Room Example

Our implementation demonstrates a chat room where:

### User Types:
- **Regular Users**: Can send messages, private messages, and change status
- **Moderator Users**: Can send announcements and have special privileges
- **Bot Users**: Automated responses to specific commands

### Communication Types:
- **Public Messages**: Broadcast to all users
- **Private Messages**: Sent to specific users
- **Announcements**: Special messages from moderators
- **Status Changes**: User status updates
- **System Events**: Join/leave notifications

## Running the Demo

```bash
cd gocode/MediatorDesignPattern
go run main.go
```

### Expected Output:
The demo shows:
- Users joining the chat room
- Various types of messages and interactions
- Bot responding to commands
- Private messaging
- Moderator announcements
- Status changes and user management
- Chat log and active user tracking

## Advantages

- **Loose Coupling**: Components don't depend on each other directly
- **Centralized Control**: All interaction logic is in one place
- **Reusability**: Components can be reused in different contexts
- **Maintainability**: Easier to modify interaction behavior
- **Single Responsibility**: Each component focuses on its core functionality

## Disadvantages

- **Complexity**: Mediator can become complex as it grows
- **Single Point of Failure**: All communication depends on the mediator
- **Performance**: Additional layer of indirection
- **God Object**: Mediator might become too large and complex

## Go-Specific Features

### Interface-Based Design
```go
type Mediator interface {
    Notify(sender Component, event string, data interface{})
    RegisterComponent(component Component)
    UnregisterComponent(component Component)
}
```

### Composition Over Inheritance
```go
type ModeratorUser struct {
    *BaseUser  // Embedded struct for composition
}
```

### Type Assertions for Event Data
```go
msgData := data.(map[string]string)
```

## Real-World Examples

- **Slack/Discord**: Chat applications with channels, users, and bots
- **Web Frameworks**: Middleware systems in HTTP frameworks
- **Game Development**: Game state management and entity communication
- **Microservices**: Service mesh and API gateways
- **UI Frameworks**: Component communication in React, Vue, etc.

## Event Types in Our Implementation

1. **send_message**: Public message broadcasting
2. **send_private_message**: Direct user-to-user messaging
3. **broadcast_announcement**: Moderator announcements
4. **request_user_list**: Request active user information
5. **change_status**: User status updates
6. **user_joined/user_left**: System events for user management

## Best Practices

1. **Keep Mediator Focused**: Don't let it become a god object
2. **Use Events**: Event-driven communication is more flexible
3. **Interface Segregation**: Keep interfaces small and focused
4. **Error Handling**: Handle cases where components aren't available
5. **State Management**: Be careful about shared state in the mediator

## Summary

The Mediator pattern provides an excellent way to manage complex interactions between multiple objects while keeping them loosely coupled. In our chat room example, users, moderators, and bots can all interact seamlessly without knowing about each other's implementation details. The pattern is particularly useful in scenarios where you have many objects that need to communicate in complex ways, such as GUI applications, chat systems, or workflow engines.

The key benefit is that it transforms a web of interconnected objects into a hub-and-spoke model, making the system easier to understand, maintain, and extend.