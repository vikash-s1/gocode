package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"SingletonDesignPattern/internal/singleton"
)

func main() {
	fmt.Println("=== Singleton Design Pattern Demo ===\n")

	// Demonstrate Database Singleton
	demonstrateDatabase()
	
	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	
	// Demonstrate Logger Singleton
	demonstrateLogger()
	
	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	
	// Demonstrate Thread Safety
	demonstrateThreadSafety()
}

func demonstrateDatabase() {
	fmt.Println("1. Database Singleton Demo:")
	
	// Get first instance
	db1 := singleton.GetInstance()
	fmt.Printf("DB1 Connection String: %s\n", db1.GetConnectionString())
	
	// Get second instance
	db2 := singleton.GetInstance()
	fmt.Printf("DB2 Connection String: %s\n", db2.GetConnectionString())
	
	// Verify they are the same instance
	fmt.Printf("Are db1 and db2 the same instance? %t\n", db1 == db2)
	
	// Use the database
	db1.Connect()
	fmt.Printf("DB2 is connected: %t\n", db2.IsConnected())
	
	db2.Disconnect()
	fmt.Printf("DB1 is connected: %t\n", db1.IsConnected())
}

func demonstrateLogger() {
	fmt.Println("2. Logger Singleton Demo:")
	
	// Get first logger instance
	logger1 := singleton.GetLogger()
	logger1.Log("INFO", "First log message")
	
	// Get second logger instance
	logger2 := singleton.GetLogger()
	logger2.Log("DEBUG", "Second log message")
	
	// Verify they are the same instance
	fmt.Printf("Are logger1 and logger2 the same instance? %t\n", logger1 == logger2)
	
	// Change log level using one instance
	logger1.SetLogLevel("ERROR")
	fmt.Printf("Logger2 log level: %s\n", logger2.GetLogLevel())
}

func demonstrateThreadSafety() {
	fmt.Println("3. Thread Safety Demo:")
	
	var wg sync.WaitGroup
	instances := make([]*singleton.DatabaseConnection, 10)
	
	// Create 10 goroutines trying to get the singleton instance
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 10) // Simulate some work
			instances[index] = singleton.GetInstance()
			fmt.Printf("Goroutine %d got instance\n", index)
		}(i)
	}
	
	wg.Wait()
	
	// Verify all instances are the same
	allSame := true
	for i := 1; i < len(instances); i++ {
		if instances[0] != instances[i] {
			allSame = false
			break
		}
	}
	
	fmt.Printf("All 10 instances are the same: %t\n", allSame)
}