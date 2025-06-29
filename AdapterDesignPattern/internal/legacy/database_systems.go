// Package legacy contains old database systems with incompatible interfaces
package legacy

import (
	"fmt"
	"time"
)

// MySQLDatabase represents a legacy MySQL database connection
type MySQLDatabase struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Database   string
	connection bool
}

// OpenConnection opens a MySQL connection using legacy method
func (m *MySQLDatabase) OpenConnection() error {
	fmt.Printf("   ✓ MySQL: Opening connection to %s:%d\n", m.Host, m.Port)
	m.connection = true
	return nil
}

// CloseConnection closes the MySQL connection
func (m *MySQLDatabase) CloseConnection() error {
	fmt.Printf("   ✓ MySQL: Closing connection\n")
	m.connection = false
	return nil
}

// RunSQL executes SQL using MySQL-specific method
func (m *MySQLDatabase) RunSQL(sql string) error {
	if !m.connection {
		return fmt.Errorf("MySQL connection not open")
	}
	fmt.Printf("   ✓ MySQL: Executing SQL: %s\n", sql)
	return nil
}

// GetMySQLInfo returns MySQL-specific connection information
func (m *MySQLDatabase) GetMySQLInfo() string {
	return fmt.Sprintf("MySQL %s:%d (Connected: %v)", m.Host, m.Port, m.connection)
}

// PostgreSQLDatabase represents a legacy PostgreSQL database connection
type PostgreSQLDatabase struct {
	ConnectionString string
	connected        bool
	lastQuery        string
}

// EstablishConnection connects to PostgreSQL using legacy method
func (p *PostgreSQLDatabase) EstablishConnection() error {
	fmt.Printf("   ✓ PostgreSQL: Establishing connection with: %s\n", p.ConnectionString)
	p.connected = true
	return nil
}

// TerminateConnection disconnects from PostgreSQL
func (p *PostgreSQLDatabase) TerminateConnection() error {
	fmt.Printf("   ✓ PostgreSQL: Terminating connection\n")
	p.connected = false
	return nil
}

// ExecuteSQL runs SQL using PostgreSQL-specific method
func (p *PostgreSQLDatabase) ExecuteSQL(query string) error {
	if !p.connected {
		return fmt.Errorf("PostgreSQL not connected")
	}
	fmt.Printf("   ✓ PostgreSQL: Executing query: %s\n", query)
	p.lastQuery = query
	return nil
}

// GetPostgreSQLStatus returns PostgreSQL-specific status
func (p *PostgreSQLDatabase) GetPostgreSQLStatus() string {
	return fmt.Sprintf("PostgreSQL (Connected: %v, Last Query: %s)", p.connected, p.lastQuery)
}

// MongoClient represents a legacy MongoDB client
type MongoClient struct {
	URI        string
	Database   string
	Collection string
	session    bool
}

// StartSession starts a MongoDB session using legacy method
func (m *MongoClient) StartSession() error {
	fmt.Printf("   ✓ MongoDB: Starting session with URI: %s\n", m.URI)
	m.session = true
	return nil
}

// EndSession ends the MongoDB session
func (m *MongoClient) EndSession() error {
	fmt.Printf("   ✓ MongoDB: Ending session\n")
	m.session = false
	return nil
}

// FindDocuments executes a find operation using MongoDB-specific method
func (m *MongoClient) FindDocuments(filter string) error {
	if !m.session {
		return fmt.Errorf("MongoDB session not started")
	}
	fmt.Printf("   ✓ MongoDB: Finding documents with filter: %s\n", filter)
	return nil
}

// GetMongoStatus returns MongoDB-specific status
func (m *MongoClient) GetMongoStatus() string {
	return fmt.Sprintf("MongoDB %s (Session Active: %v)", m.URI, m.session)
}