package singleton

import (
	"fmt"
	"sync"
)

// DatabaseConnection represents a singleton database connection
type DatabaseConnection struct {
	connectionString string
	isConnected      bool
}

var (
	// instance holds the single instance of DatabaseConnection
	instance *DatabaseConnection
	// once ensures the instance is created only once
	once sync.Once
)

// GetInstance returns the singleton instance of DatabaseConnection
// This method is thread-safe using sync.Once
func GetInstance() *DatabaseConnection {
	once.Do(func() {
		instance = &DatabaseConnection{
			connectionString: "localhost:5432/mydb",
			isConnected:      false,
		}
		fmt.Println("Creating new database connection instance")
	})
	return instance
}

// Connect establishes the database connection
func (db *DatabaseConnection) Connect() {
	if !db.isConnected {
		db.isConnected = true
		fmt.Printf("Connected to database: %s\n", db.connectionString)
	} else {
		fmt.Println("Already connected to database")
	}
}

// Disconnect closes the database connection
func (db *DatabaseConnection) Disconnect() {
	if db.isConnected {
		db.isConnected = false
		fmt.Println("Disconnected from database")
	} else {
		fmt.Println("Database is not connected")
	}
}

// GetConnectionString returns the connection string
func (db *DatabaseConnection) GetConnectionString() string {
	return db.connectionString
}

// IsConnected returns the connection status
func (db *DatabaseConnection) IsConnected() bool {
	return db.isConnected
}