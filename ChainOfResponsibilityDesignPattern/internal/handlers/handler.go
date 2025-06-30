package handlers

// Request represents a request that needs to be processed
type Request struct {
	Type    string
	Content string
	Level   int
}

// Handler defines the interface for all handlers in the chain
type Handler interface {
	SetNext(handler Handler)
	Handle(request *Request) bool
}