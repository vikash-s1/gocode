package handlers

import "fmt"

// BaseHandler provides common functionality for all handlers
type BaseHandler struct {
	next Handler
}

// SetNext sets the next handler in the chain
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// HandleNext passes the request to the next handler if it exists
func (h *BaseHandler) HandleNext(request *Request) bool {
	if h.next != nil {
		return h.next.Handle(request)
	}
	fmt.Printf("No handler found for request: %s\n", request.Type)
	return false
}