package handlers

import "fmt"

// InfoHandler handles informational requests
type InfoHandler struct {
	BaseHandler
}

// NewInfoHandler creates a new InfoHandler
func NewInfoHandler() *InfoHandler {
	return &InfoHandler{}
}

// Handle processes info-level requests
func (h *InfoHandler) Handle(request *Request) bool {
	if request.Type == "INFO" || request.Level <= 1 {
		fmt.Printf("[INFO] Processing request: %s\n", request.Content)
		return true
	}
	return h.HandleNext(request)
}