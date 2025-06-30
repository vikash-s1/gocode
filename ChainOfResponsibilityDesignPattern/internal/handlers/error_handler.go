package handlers

import "fmt"

// ErrorHandler handles error requests
type ErrorHandler struct {
	BaseHandler
}

// NewErrorHandler creates a new ErrorHandler
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

// Handle processes error-level requests
func (h *ErrorHandler) Handle(request *Request) bool {
	if request.Type == "ERROR" || request.Level >= 3 {
		fmt.Printf("[ERROR] Processing critical request: %s\n", request.Content)
		return true
	}
	return h.HandleNext(request)
}