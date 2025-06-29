package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"statepattern/internal/states"
)

func main() {
	fmt.Println("🎰 Welcome to the State Pattern Vending Machine Demo!")
	fmt.Println(strings.Repeat("=", 50))
	
	// Create a vending machine with 3 products
	vendingMachine := states.NewVendingMachine(3)
	
	// Interactive demo
	runInteractiveDemo(vendingMachine)
	
	fmt.Println("\n🎯 Automated Demo:")
	fmt.Println(strings.Repeat("=", 30))
	
	// Automated demo to show all state transitions
	runAutomatedDemo()
}

func runInteractiveDemo(vm *states.VendingMachine) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("\n📊 Current State: %s | Products: %d\n", vm.GetCurrentState(), vm.GetProductCount())
	fmt.Println("\nCommands:")
	fmt.Println("1 - Insert Coin")
	fmt.Println("2 - Select Product") 
	fmt.Println("3 - Dispense Product")
	fmt.Println("4 - Cancel Transaction")
	fmt.Println("5 - Check Status")
	fmt.Println("q - Quit to Automated Demo")
	
	for {
		fmt.Print("\nEnter command: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		
		switch input {
		case "1":
			vm.InsertCoin()
		case "2":
			vm.SelectProduct()
		case "3":
			vm.DispenseProduct()
		case "4":
			vm.Cancel()
		case "5":
			fmt.Printf("📊 Current State: %s | Products: %d | Coin: %v\n", 
				vm.GetCurrentState(), vm.GetProductCount(), vm.IsCoinInserted())
		case "q":
			return
		default:
			fmt.Println("❌ Invalid command. Try again.")
			continue
		}
		
		fmt.Printf("📊 Current State: %s | Products: %d\n", vm.GetCurrentState(), vm.GetProductCount())
		
		if vm.GetCurrentState() == "OutOfStock" {
			fmt.Println("🔚 Machine is out of stock. Moving to automated demo...")
			return
		}
	}
}

func runAutomatedDemo() {
	// Demo 1: Normal flow
	fmt.Println("\n🎬 Demo 1: Normal Purchase Flow")
	fmt.Println(strings.Repeat("-", 35))
	vm1 := states.NewVendingMachine(2)
	
	demonstrateFlow(vm1, []string{
		"Insert Coin",
		"Select Product", 
		"Dispense Product",
	})
	
	// Demo 2: Cancellation flow
	fmt.Println("\n🎬 Demo 2: Cancellation Flow")
	fmt.Println(strings.Repeat("-", 30))
	vm2 := states.NewVendingMachine(2)
	
	demonstrateFlow(vm2, []string{
		"Insert Coin",
		"Cancel Transaction",
	})
	
	// Demo 3: Invalid operations
	fmt.Println("\n🎬 Demo 3: Invalid Operations")
	fmt.Println(strings.Repeat("-", 32))
	vm3 := states.NewVendingMachine(2)
	
	demonstrateFlow(vm3, []string{
		"Select Product (without coin)",
		"Dispense Product (without selection)",
		"Insert Coin",
		"Insert Coin (again)",
		"Select Product",
		"Cancel (during dispensing)",
		"Dispense Product",
	})
	
	// Demo 4: Out of stock scenario
	fmt.Println("\n🎬 Demo 4: Out of Stock Scenario")
	fmt.Println(strings.Repeat("-", 33))
	vm4 := states.NewVendingMachine(1) // Only 1 product
	
	demonstrateFlow(vm4, []string{
		"Insert Coin",
		"Select Product",
		"Dispense Product", // This will make it out of stock
		"Insert Coin (out of stock)",
	})
}

func demonstrateFlow(vm *states.VendingMachine, actions []string) {
	for i, action := range actions {
		fmt.Printf("\n%d. %s\n", i+1, action)
		fmt.Printf("   State before: %s\n", vm.GetCurrentState())
		
		switch {
		case strings.Contains(action, "Insert Coin"):
			vm.InsertCoin()
		case strings.Contains(action, "Select Product"):
			vm.SelectProduct()
		case strings.Contains(action, "Dispense Product"):
			vm.DispenseProduct()
		case strings.Contains(action, "Cancel"):
			vm.Cancel()
		}
		
		fmt.Printf("   State after:  %s\n", vm.GetCurrentState())
		fmt.Printf("   Products left: %d\n", vm.GetProductCount())
	}
}