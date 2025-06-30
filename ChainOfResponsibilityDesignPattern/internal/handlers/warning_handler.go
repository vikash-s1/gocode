package handlers

import "fmt"

// WarningHandler handles warning requests
type WarningHandler struct {
	BaseHandler
}

// NewWarningHandler creates a new WarningHandler
func NewWarningHandler() *WarningHandler {
	return &WarningHandler{}
}

// Handle processes warning-level requests
func (h *WarningHandler) Handle(request *Request) bool {
	if request.Type == "WARNING" || request.Level == 2 {
		fmt.Printf("[WARNING] Processing request: %s\n", request.Content)
		return true
	}
	return h.HandleNext(request)
}