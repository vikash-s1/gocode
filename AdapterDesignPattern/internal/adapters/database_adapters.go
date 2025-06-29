// Package adapters contains database adapter implementations
package adapters

import (
	"fmt"
	"strings"

	"github.com/example/adapter-pattern/internal/legacy"
	"github.com/example/adapter-pattern/internal/modern"
)

// MySQLAdapter adapts the legacy MySQL database to the modern database interface
type MySQLAdapter struct {
	database *legacy.MySQLDatabase
}

// NewMySQLAdapter creates a new MySQL adapter
func NewMySQLAdapter(database *legacy.MySQLDatabase) *MySQLAdapter {
	return &MySQLAdapter{database: database}
}

// Connect adapts MySQL's OpenConnection method to our standard interface
func (m *MySQLAdapter) Connect() error {
	return m.database.OpenConnection()
}

// Disconnect adapts MySQL's CloseConnection method to our standard interface
func (m *MySQLAdapter) Disconnect() error {
	return m.database.CloseConnection()
}

// ExecuteQuery adapts MySQL's RunSQL method to our standard interface
func (m *MySQLAdapter) ExecuteQuery(query string) error {
	return m.database.RunSQL(query)
}

// ExecuteTransaction executes multiple queries in a transaction (simulated for MySQL)
func (m *MySQLAdapter) ExecuteTransaction(queries []string) error {
	// Start transaction
	if err := m.database.RunSQL("START TRANSACTION"); err != nil {
		return err
	}
	
	// Execute all queries
	for _, query := range queries {
		if err := m.database.RunSQL(query); err != nil {
			// Rollback on error
			m.database.RunSQL("ROLLBACK")
			return fmt.Errorf("transaction failed: %w", err)
		}
	}
	
	// Commit transaction
	return m.database.RunSQL("COMMIT")
}

// GetConnectionInfo adapts MySQL's info method to our standard interface
func (m *MySQLAdapter) GetConnectionInfo() string {
	return m.database.GetMySQLInfo()
}

// PostgreSQLAdapter adapts the legacy PostgreSQL database to the modern database interface
type PostgreSQLAdapter struct {
	database *legacy.PostgreSQLDatabase
}

// NewPostgreSQLAdapter creates a new PostgreSQL adapter
func NewPostgreSQLAdapter(database *legacy.PostgreSQLDatabase) *PostgreSQLAdapter {
	return &PostgreSQLAdapter{database: database}
}

// Connect adapts PostgreSQL's EstablishConnection method to our standard interface
func (p *PostgreSQLAdapter) Connect() error {
	return p.database.EstablishConnection()
}

// Disconnect adapts PostgreSQL's TerminateConnection method to our standard interface
func (p *PostgreSQLAdapter) Disconnect() error {
	return p.database.TerminateConnection()
}

// ExecuteQuery adapts PostgreSQL's ExecuteSQL method to our standard interface
func (p *PostgreSQLAdapter) ExecuteQuery(query string) error {
	return p.database.ExecuteSQL(query)
}

// ExecuteTransaction executes multiple queries in a transaction (simulated for PostgreSQL)
func (p *PostgreSQLAdapter) ExecuteTransaction(queries []string) error {
	// Start transaction
	if err := p.database.ExecuteSQL("BEGIN"); err != nil {
		return err
	}
	
	// Execute all queries
	for _, query := range queries {
		if err := p.database.ExecuteSQL(query); err != nil {
			// Rollback on error
			p.database.ExecuteSQL("ROLLBACK")
			return fmt.Errorf("transaction failed: %w", err)
		}
	}
	
	// Commit transaction
	return p.database.ExecuteSQL("COMMIT")
}

// GetConnectionInfo adapts PostgreSQL's status method to our standard interface
func (p *PostgreSQLAdapter) GetConnectionInfo() string {
	return p.database.GetPostgreSQLStatus()
}

// MongoAdapter adapts the legacy MongoDB client to the modern database interface
type MongoAdapter struct {
	client *legacy.MongoClient
}

// NewMongoAdapter creates a new MongoDB adapter
func NewMongoAdapter(client *legacy.MongoClient) *MongoAdapter {
	return &MongoAdapter{client: client}
}

// Connect adapts MongoDB's StartSession method to our standard interface
func (m *MongoAdapter) Connect() error {
	return m.client.StartSession()
}

// Disconnect adapts MongoDB's EndSession method to our standard interface
func (m *MongoAdapter) Disconnect() error {
	return m.client.EndSession()
}

// ExecuteQuery adapts MongoDB's FindDocuments method to our standard interface
func (m *MongoAdapter) ExecuteQuery(query string) error {
	// Convert SQL-like query to MongoDB filter (simplified)
	mongoFilter := m.convertSQLToMongoFilter(query)
	return m.client.FindDocuments(mongoFilter)
}

// ExecuteTransaction executes multiple operations in a transaction (simulated for MongoDB)
func (m *MongoAdapter) ExecuteTransaction(queries []string) error {
	// MongoDB transactions are more complex, this is a simplified simulation
	fmt.Printf("   ✓ MongoDB: Starting transaction with %d operations\n", len(queries))
	
	for i, query := range queries {
		mongoFilter := m.convertSQLToMongoFilter(query)
		if err := m.client.FindDocuments(fmt.Sprintf("Transaction[%d]: %s", i+1, mongoFilter)); err != nil {
			return fmt.Errorf("transaction operation %d failed: %w", i+1, err)
		}
	}
	
	fmt.Printf("   ✓ MongoDB: Transaction completed successfully\n")
	return nil
}

// GetConnectionInfo adapts MongoDB's status method to our standard interface
func (m *MongoAdapter) GetConnectionInfo() string {
	return m.client.GetMongoStatus()
}

// convertSQLToMongoFilter converts a simple SQL query to MongoDB filter (simplified)
func (m *MongoAdapter) convertSQLToMongoFilter(sqlQuery string) string {
	// This is a very simplified conversion for demonstration
	if strings.Contains(strings.ToUpper(sqlQuery), "SELECT * FROM USERS") {
		return "{}"
	}
	if strings.Contains(strings.ToUpper(sqlQuery), "WHERE") {
		return "{status: 'active'}"
	}
	return "{query: 'converted'}"
}

// DatabaseAdapterFactory creates database adapters based on database type
type DatabaseAdapterFactory struct{}

// CreateDatabaseAdapter creates the appropriate database adapter based on database type
func (f *DatabaseAdapterFactory) CreateDatabaseAdapter(dbType string, config map[string]string) (modern.Database, error) {
	switch strings.ToLower(dbType) {
	case "mysql":
		mysqlDB := &legacy.MySQLDatabase{
			Host: config["host"],
			Port: 3306, // default port
		}
		return NewMySQLAdapter(mysqlDB), nil
		
	case "postgresql", "postgres":
		postgresDB := &legacy.PostgreSQLDatabase{
			ConnectionString: config["connection_string"],
		}
		return NewPostgreSQLAdapter(postgresDB), nil
		
	case "mongodb", "mongo":
		mongoClient := &legacy.MongoClient{
			URI: config["uri"],
		}
		return NewMongoAdapter(mongoClient), nil
		
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}