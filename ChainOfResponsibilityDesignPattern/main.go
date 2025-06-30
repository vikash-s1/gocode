package main

import (
	"chainofresponsibility/internal/handlers"
	"fmt"
)

func main() {
	fmt.Println("=== Chain of Responsibility Design Pattern Demo ===\n")

	// Create handlers
	infoHandler := handlers.NewInfoHandler()
	warningHandler := handlers.NewWarningHandler()
	errorHandler := handlers.NewErrorHandler()

	// Build the chain: Info -> Warning -> Error
	infoHandler.SetNext(warningHandler)
	warningHandler.SetNext(errorHandler)

	// Create test requests
	requests := []*handlers.Request{
		{Type: "INFO", Content: "System startup completed", Level: 1},
		{Type: "WARNING", Content: "Memory usage is high", Level: 2},
		{Type: "ERROR", Content: "Database connection failed", Level: 3},
		{Type: "UNKNOWN", Content: "Unknown request type", Level: 0},
		{Type: "INFO", Content: "User logged in", Level: 1},
		{Type: "ERROR", Content: "Critical system failure", Level: 4},
	}

	// Process requests through the chain
	fmt.Println("Processing requests through the chain:\n")
	for i, request := range requests {
		fmt.Printf("Request %d:\n", i+1)
		handled := infoHandler.Handle(request)
		if !handled {
			fmt.Printf("Request was not handled by any handler\n")
		}
		fmt.Println()
	}

	fmt.Println("=== Alternative Chain Configuration ===\n")
	
	// Create a different chain: Error -> Warning -> Info
	errorHandler2 := handlers.NewErrorHandler()
	warningHandler2 := handlers.NewWarningHandler()
	infoHandler2 := handlers.NewInfoHandler()
	
	errorHandler2.SetNext(warningHandler2)
	warningHandler2.SetNext(infoHandler2)

	fmt.Println("Processing same requests with different chain order:\n")
	for i, request := range requests[:3] {
		fmt.Printf("Request %d:\n", i+1)
		errorHandler2.Handle(request)
		fmt.Println()
	}
}