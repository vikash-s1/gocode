package builder

import (
	"errors"
	"fmt"
)

// Computer represents the complex product being built
type Computer struct {
	CPU        string
	RAM        int // in GB
	Storage    string
	GPU        string
	OS         string
	Monitor    string
	Keyboard   string
	Mouse      string
	PowerSupply int // in watts
}

// String returns a formatted string representation of the computer
func (c *Computer) String() string {
	return fmt.Sprintf(`Computer Specifications:
- CPU: %s
- RAM: %d GB
- Storage: %s
- GPU: %s
- OS: %s
- Monitor: %s
- Keyboard: %s
- Mouse: %s
- Power Supply: %d watts`,
		c.CPU, c.RAM, c.Storage, c.GPU, c.OS,
		c.Monitor, c.Keyboard, c.Mouse, c.PowerSupply)
}

// ComputerBuilder defines the builder interface
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram int) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	SetOS(os string) ComputerBuilder
	SetMonitor(monitor string) ComputerBuilder
	SetKeyboard(keyboard string) ComputerBuilder
	SetMouse(mouse string) ComputerBuilder
	SetPowerSupply(watts int) ComputerBuilder
	Build() (*Computer, error)
}

// computerBuilder is the concrete implementation of ComputerBuilder
type computerBuilder struct {
	computer *Computer
}

// NewComputerBuilder creates a new computer builder instance
func NewComputerBuilder() ComputerBuilder {
	return &computerBuilder{
		computer: &Computer{},
	}
}

// SetCPU sets the CPU specification
func (cb *computerBuilder) SetCPU(cpu string) ComputerBuilder {
	cb.computer.CPU = cpu
	return cb
}

// SetRAM sets the RAM amount in GB
func (cb *computerBuilder) SetRAM(ram int) ComputerBuilder {
	cb.computer.RAM = ram
	return cb
}

// SetStorage sets the storage specification
func (cb *computerBuilder) SetStorage(storage string) ComputerBuilder {
	cb.computer.Storage = storage
	return cb
}

// SetGPU sets the GPU specification
func (cb *computerBuilder) SetGPU(gpu string) ComputerBuilder {
	cb.computer.GPU = gpu
	return cb
}

// SetOS sets the operating system
func (cb *computerBuilder) SetOS(os string) ComputerBuilder {
	cb.computer.OS = os
	return cb
}

// SetMonitor sets the monitor specification
func (cb *computerBuilder) SetMonitor(monitor string) ComputerBuilder {
	cb.computer.Monitor = monitor
	return cb
}

// SetKeyboard sets the keyboard specification
func (cb *computerBuilder) SetKeyboard(keyboard string) ComputerBuilder {
	cb.computer.Keyboard = keyboard
	return cb
}

// SetMouse sets the mouse specification
func (cb *computerBuilder) SetMouse(mouse string) ComputerBuilder {
	cb.computer.Mouse = mouse
	return cb
}

// SetPowerSupply sets the power supply wattage
func (cb *computerBuilder) SetPowerSupply(watts int) ComputerBuilder {
	cb.computer.PowerSupply = watts
	return cb
}

// Build constructs and validates the final computer object
func (cb *computerBuilder) Build() (*Computer, error) {
	// Validation logic
	if cb.computer.CPU == "" {
		return nil, errors.New("CPU is required")
	}
	if cb.computer.RAM <= 0 {
		return nil, errors.New("RAM must be greater than 0")
	}
	if cb.computer.Storage == "" {
		return nil, errors.New("Storage is required")
	}
	if cb.computer.OS == "" {
		return nil, errors.New("Operating System is required")
	}
	if cb.computer.PowerSupply <= 0 {
		return nil, errors.New("Power Supply must be greater than 0 watts")
	}

	// Create a copy to ensure immutability
	result := &Computer{
		CPU:         cb.computer.CPU,
		RAM:         cb.computer.RAM,
		Storage:     cb.computer.Storage,
		GPU:         cb.computer.GPU,
		OS:          cb.computer.OS,
		Monitor:     cb.computer.Monitor,
		Keyboard:    cb.computer.Keyboard,
		Mouse:       cb.computer.Mouse,
		PowerSupply: cb.computer.PowerSupply,
	}

	return result, nil
}