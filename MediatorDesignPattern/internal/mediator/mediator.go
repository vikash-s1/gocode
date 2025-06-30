package mediator

// Mediator defines the interface for communication between components
type Mediator interface {
	Notify(sender Component, event string, data interface{})
	RegisterComponent(component Component)
	UnregisterComponent(component Component)
}

// Component defines the interface for components that communicate through mediator
type Component interface {
	SetMediator(mediator Mediator)
	GetName() string
	HandleNotification(event string, data interface{})
}