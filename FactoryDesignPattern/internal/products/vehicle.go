package products

import "fmt"

// Vehicle interface defines the contract for all vehicle products
type Vehicle interface {
	Start() string
	Stop() string
	GetInfo() string
}

// Car represents a concrete car product
type Car struct {
	Brand string
	Model string
}

func (c *Car) Start() string {
	return fmt.Sprintf("%s %s engine started", c.Brand, c.Model)
}

func (c *Car) Stop() string {
	return fmt.Sprintf("%s %s engine stopped", c.Brand, c.Model)
}

func (c *Car) GetInfo() string {
	return fmt.Sprintf("Car: %s %s", c.Brand, c.Model)
}

// Motorcycle represents a concrete motorcycle product
type Motorcycle struct {
	Brand string
	Model string
}

func (m *Motorcycle) Start() string {
	return fmt.Sprintf("%s %s motorcycle started", m.Brand, m.Model)
}

func (m *Motorcycle) Stop() string {
	return fmt.Sprintf("%s %s motorcycle stopped", m.Brand, m.Model)
}

func (m *Motorcycle) GetInfo() string {
	return fmt.Sprintf("Motorcycle: %s %s", m.Brand, m.Model)
}

// Truck represents a concrete truck product
type Truck struct {
	Brand    string
	Model    string
	Capacity int
}

func (t *Truck) Start() string {
	return fmt.Sprintf("%s %s truck engine started", t.Brand, t.Model)
}

func (t *Truck) Stop() string {
	return fmt.Sprintf("%s %s truck engine stopped", t.Brand, t.Model)
}

func (t *Truck) GetInfo() string {
	return fmt.Sprintf("Truck: %s %s (Capacity: %d tons)", t.Brand, t.Model, t.Capacity)
}