package pool

import (
	"fmt"
	"time"
)

// Connection represents a database connection
type Connection struct {
	ID        int
	InUse     bool
	CreatedAt time.Time
}

// NewConnection creates a new database connection
func NewConnection(id int) *Connection {
	fmt.Printf("Creating new connection with ID: %d\n", id)
	return &Connection{
		ID:        id,
		InUse:     false,
		CreatedAt: time.Now(),
	}
}

// Connect simulates establishing a connection
func (c *Connection) Connect() {
	fmt.Printf("Connection %d: Establishing connection...\n", c.ID)
	c.InUse = true
}

// Disconnect simulates closing a connection
func (c *Connection) Disconnect() {
	fmt.Printf("Connection %d: Closing connection...\n", c.ID)
	c.InUse = false
}

// Execute simulates executing a query
func (c *Connection) Execute(query string) {
	if !c.InUse {
		fmt.Printf("Connection %d: Cannot execute - connection not active\n", c.ID)
		return
	}
	fmt.Printf("Connection %d: Executing query: %s\n", c.ID, query)
}