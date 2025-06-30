package director

import (
	"builderpattern/internal/builder"
)

// ComputerDirector orchestrates the construction of different types of computers
type ComputerDirector struct {
	builder builder.ComputerBuilder
}

// NewComputerDirector creates a new computer director
func NewComputerDirector(builder builder.ComputerBuilder) *ComputerDirector {
	return &ComputerDirector{
		builder: builder,
	}
}

// BuildGamingComputer constructs a high-end gaming computer
func (cd *ComputerDirector) BuildGamingComputer() (*builder.Computer, error) {
	return cd.builder.
		SetCPU("Intel Core i9-13900K").
		SetRAM(32).
		SetStorage("2TB NVMe SSD").
		SetGPU("NVIDIA RTX 4080").
		SetOS("Windows 11 Pro").
		SetMonitor("27\" 4K Gaming Monitor").
		SetKeyboard("Mechanical RGB Gaming Keyboard").
		SetMouse("High-DPI Gaming Mouse").
		SetPowerSupply(850).
		Build()
}

// BuildOfficeComputer constructs a basic office computer
func (cd *ComputerDirector) BuildOfficeComputer() (*builder.Computer, error) {
	return cd.builder.
		SetCPU("Intel Core i5-12400").
		SetRAM(16).
		SetStorage("512GB SSD").
		SetGPU("Integrated Graphics").
		SetOS("Windows 11").
		SetMonitor("24\" 1080p Monitor").
		SetKeyboard("Standard Keyboard").
		SetMouse("Optical Mouse").
		SetPowerSupply(450).
		Build()
}

// BuildWorkstationComputer constructs a professional workstation
func (cd *ComputerDirector) BuildWorkstationComputer() (*builder.Computer, error) {
	return cd.builder.
		SetCPU("AMD Ryzen 9 7950X").
		SetRAM(64).
		SetStorage("4TB NVMe SSD").
		SetGPU("NVIDIA RTX A4000").
		SetOS("Ubuntu 22.04 LTS").
		SetMonitor("32\" 4K Professional Monitor").
		SetKeyboard("Ergonomic Keyboard").
		SetMouse("Precision Mouse").
		SetPowerSupply(1000).
		Build()
}

// BuildBudgetComputer constructs a budget-friendly computer
func (cd *ComputerDirector) BuildBudgetComputer() (*builder.Computer, error) {
	return cd.builder.
		SetCPU("AMD Ryzen 5 5600G").
		SetRAM(8).
		SetStorage("256GB SSD").
		SetGPU("Integrated Graphics").
		SetOS("Linux Mint").
		SetMonitor("21\" 1080p Monitor").
		SetKeyboard("Basic Keyboard").
		SetMouse("Basic Mouse").
		SetPowerSupply(350).
		Build()
}