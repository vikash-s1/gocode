package main

import (
	"fmt"
	"objectpool/internal/pool"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Object Pool Design Pattern Demo ===\n")
	
	// Create a connection pool with max 3 connections
	connectionPool := pool.NewConnectionPool(3)
	defer connectionPool.Close()
	
	// Demo 1: Basic usage
	fmt.Println("--- Demo 1: Basic Connection Pool Usage ---")
	basicUsageDemo(connectionPool)
	
	// Demo 2: Pool exhaustion
	fmt.Println("\n--- Demo 2: Pool Exhaustion Scenario ---")
	poolExhaustionDemo(connectionPool)
	
	// Demo 3: Concurrent access
	fmt.Println("\n--- Demo 3: Concurrent Access ---")
	concurrentAccessDemo(connectionPool)
	
	// Demo 4: Pool statistics
	fmt.Println("\n--- Demo 4: Pool Statistics ---")
	poolStatsDemo(connectionPool)
}

func basicUsageDemo(cp *pool.ConnectionPool) {
	// Get a connection
	conn1, err := cp.GetConnection()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Use the connection
	conn1.Execute("SELECT * FROM users")
	
	// Get another connection
	conn2, err := cp.GetConnection()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	conn2.Execute("SELECT * FROM products")
	
	// Release connections back to pool
	cp.ReleaseConnection(conn1)
	cp.ReleaseConnection(conn2)
	
	// Reuse connection from pool
	conn3, err := cp.GetConnection()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	conn3.Execute("UPDATE users SET status = 'active'")
	cp.ReleaseConnection(conn3)
}

func poolExhaustionDemo(cp *pool.ConnectionPool) {
	var connections []*pool.Connection
	
	// Try to get more connections than pool limit
	for i := 0; i < 5; i++ {
		conn, err := cp.GetConnection()
		if err != nil {
			fmt.Printf("Failed to get connection %d: %v\n", i+1, err)
			break
		}
		connections = append(connections, conn)
		conn.Execute(fmt.Sprintf("Query %d", i+1))
	}
	
	// Release all connections
	for _, conn := range connections {
		cp.ReleaseConnection(conn)
	}
}

func concurrentAccessDemo(cp *pool.ConnectionPool) {
	var wg sync.WaitGroup
	
	// Simulate 5 concurrent workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			conn, err := cp.GetConnection()
			if err != nil {
				fmt.Printf("Worker %d: Failed to get connection: %v\n", workerID, err)
				return
			}
			
			// Simulate work
			conn.Execute(fmt.Sprintf("Worker %d query", workerID))
			time.Sleep(100 * time.Millisecond)
			
			cp.ReleaseConnection(conn)
			fmt.Printf("Worker %d: Completed\n", workerID)
		}(i + 1)
	}
	
	wg.Wait()
	fmt.Println("All workers completed")
}

func poolStatsDemo(cp *pool.ConnectionPool) {
	available, total := cp.GetPoolStats()
	fmt.Printf("Pool Stats - Available: %d, Total Created: %d\n", available, total)
	
	// Get some connections
	conn1, _ := cp.GetConnection()
	conn2, _ := cp.GetConnection()
	
	available, total = cp.GetPoolStats()
	fmt.Printf("After getting 2 connections - Available: %d, Total Created: %d\n", available, total)
	
	// Release one connection
	cp.ReleaseConnection(conn1)
	
	available, total = cp.GetPoolStats()
	fmt.Printf("After releasing 1 connection - Available: %d, Total Created: %d\n", available, total)
	
	// Release the other connection
	cp.ReleaseConnection(conn2)
	
	available, total = cp.GetPoolStats()
	fmt.Printf("After releasing all connections - Available: %d, Total Created: %d\n", available, total)
}