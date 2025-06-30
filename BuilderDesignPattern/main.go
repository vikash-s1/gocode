package main

import (
	"builderpattern/internal/builder"
	"builderpattern/internal/director"
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("=== Builder Design Pattern Demo ===\n")

	// Example 1: Direct builder usage
	fmt.Println("1. Building a Custom Computer using Builder directly:")
	fmt.Println(strings.Repeat("-", 50))
	
	customComputer, err := builder.NewComputerBuilder().
		SetCPU("Intel Core i7-13700K").
		SetRAM(16).
		SetStorage("1TB NVMe SSD").
		SetGPU("NVIDIA RTX 4070").
		SetOS("Windows 11").
		SetMonitor("27\" 1440p Monitor").
		SetKeyboard("Mechanical Keyboard").
		SetMouse("Gaming Mouse").
		SetPowerSupply(650).
		Build()

	if err != nil {
		log.Printf("Error building custom computer: %v", err)
	} else {
		fmt.Println(customComputer)
	}

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// Example 2: Using Director for predefined configurations
	fmt.Println("2. Using Director for Predefined Configurations:")
	fmt.Println(strings.Repeat("-", 50))

	// Gaming Computer
	fmt.Println("🎮 Gaming Computer:")
	gamingBuilder := builder.NewComputerBuilder()
	director := director.NewComputerDirector(gamingBuilder)
	
	gamingComputer, err := director.BuildGamingComputer()
	if err != nil {
		log.Printf("Error building gaming computer: %v", err)
	} else {
		fmt.Println(gamingComputer)
	}

	fmt.Println("\n" + strings.Repeat("-", 30) + "\n")

	// Office Computer
	fmt.Println("💼 Office Computer:")
	officeBuilder := builder.NewComputerBuilder()
	director = director.NewComputerDirector(officeBuilder)
	
	officeComputer, err := director.BuildOfficeComputer()
	if err != nil {
		log.Printf("Error building office computer: %v", err)
	} else {
		fmt.Println(officeComputer)
	}

	fmt.Println("\n" + strings.Repeat("-", 30) + "\n")

	// Workstation Computer
	fmt.Println("🖥️  Workstation Computer:")
	workstationBuilder := builder.NewComputerBuilder()
	director = director.NewComputerDirector(workstationBuilder)
	
	workstationComputer, err := director.BuildWorkstationComputer()
	if err != nil {
		log.Printf("Error building workstation computer: %v", err)
	} else {
		fmt.Println(workstationComputer)
	}

	fmt.Println("\n" + strings.Repeat("-", 30) + "\n")

	// Budget Computer
	fmt.Println("💰 Budget Computer:")
	budgetBuilder := builder.NewComputerBuilder()
	director = director.NewComputerDirector(budgetBuilder)
	
	budgetComputer, err := director.BuildBudgetComputer()
	if err != nil {
		log.Printf("Error building budget computer: %v", err)
	} else {
		fmt.Println(budgetComputer)
	}

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// Example 3: Error handling demonstration
	fmt.Println("3. Error Handling Demo:")
	fmt.Println(strings.Repeat("-", 50))

	// Try to build an invalid computer (missing required fields)
	invalidComputer, err := builder.NewComputerBuilder().
		SetRAM(8).
		SetStorage("500GB SSD").
		Build() // Missing CPU, OS, and PowerSupply

	if err != nil {
		fmt.Printf("❌ Expected error: %v\n", err)
	} else {
		fmt.Println("This shouldn't happen - validation failed!")
		fmt.Println(invalidComputer)
	}

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// Example 4: Minimal valid computer
	fmt.Println("4. Minimal Valid Computer:")
	fmt.Println(strings.Repeat("-", 50))

	minimalComputer, err := builder.NewComputerBuilder().
		SetCPU("Basic Processor").
		SetRAM(4).
		SetStorage("128GB SSD").
		SetOS("Linux").
		SetPowerSupply(300).
		Build()

	if err != nil {
		log.Printf("Error building minimal computer: %v", err)
	} else {
		fmt.Println(minimalComputer)
	}

	fmt.Println("\n=== Demo Complete ===")
}

