package pool

import (
	"errors"
	"fmt"
	"sync"
)

// ConnectionPool manages a pool of database connections
type ConnectionPool struct {
	pool     chan *Connection
	maxSize  int
	current  int
	mutex    sync.Mutex
	idCounter int
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool(maxSize int) *ConnectionPool {
	return &ConnectionPool{
		pool:    make(chan *Connection, maxSize),
		maxSize: maxSize,
		current: 0,
	}
}

// GetConnection retrieves a connection from the pool
func (cp *ConnectionPool) GetConnection() (*Connection, error) {
	select {
	case conn := <-cp.pool:
		// Reuse existing connection from pool
		fmt.Printf("Reusing connection %d from pool\n", conn.ID)
		conn.Connect()
		return conn, nil
	default:
		// Create new connection if pool is empty and under limit
		cp.mutex.Lock()
		defer cp.mutex.Unlock()
		
		if cp.current < cp.maxSize {
			cp.idCounter++
			conn := NewConnection(cp.idCounter)
			cp.current++
			conn.Connect()
			fmt.Printf("Pool size: %d/%d\n", cp.current, cp.maxSize)
			return conn, nil
		}
		
		return nil, errors.New("connection pool exhausted")
	}
}

// ReleaseConnection returns a connection to the pool
func (cp *ConnectionPool) ReleaseConnection(conn *Connection) {
	if conn == nil {
		return
	}
	
	conn.Disconnect()
	
	select {
	case cp.pool <- conn:
		fmt.Printf("Connection %d returned to pool\n", conn.ID)
	default:
		// Pool is full, discard the connection
		cp.mutex.Lock()
		cp.current--
		cp.mutex.Unlock()
		fmt.Printf("Connection %d discarded (pool full)\n", conn.ID)
	}
}

// GetPoolStats returns current pool statistics
func (cp *ConnectionPool) GetPoolStats() (available, total int) {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()
	return len(cp.pool), cp.current
}

// Close closes all connections in the pool
func (cp *ConnectionPool) Close() {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()
	
	close(cp.pool)
	for conn := range cp.pool {
		conn.Disconnect()
	}
	fmt.Println("Connection pool closed")
}