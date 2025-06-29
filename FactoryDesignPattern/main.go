package main

import (
	"fmt"
	"log"

	"factory-design-pattern/internal/products"
)

func main() {
	fmt.Println("=== Factory Design Pattern Demo ===\n")

	// Create a vehicle factory
	factory := products.NewVehicleFactory()

	// Demonstrate creating different types of vehicles
	demonstrateVehicleCreation(factory)

	// Demonstrate error handling
	demonstrateErrorHandling(factory)

	// Demonstrate supported types
	demonstrateSupportedTypes(factory)

	// Demonstrate truck with custom capacity
	demonstrateCustomTruck(factory)
}

func demonstrateVehicleCreation(factory *products.VehicleFactory) {
	fmt.Println("1. Creating Different Vehicle Types:")
	fmt.Println("-----------------------------------")

	// Create a car
	car, err := factory.CreateVehicle(products.CarType, "Toyota", "Camry")
	if err != nil {
		log.Printf("Error creating car: %v", err)
		return
	}
	
	fmt.Printf("✓ %s\n", car.GetInfo())
	fmt.Printf("  %s\n", car.Start())
	fmt.Printf("  %s\n", car.Stop())
	fmt.Println()

	// Create a motorcycle
	motorcycle, err := factory.CreateVehicle(products.MotorcycleType, "Harley-Davidson", "Street 750")
	if err != nil {
		log.Printf("Error creating motorcycle: %v", err)
		return
	}
	
	fmt.Printf("✓ %s\n", motorcycle.GetInfo())
	fmt.Printf("  %s\n", motorcycle.Start())
	fmt.Printf("  %s\n", motorcycle.Stop())
	fmt.Println()

	// Create a truck
	truck, err := factory.CreateVehicle(products.TruckType, "Ford", "F-150")
	if err != nil {
		log.Printf("Error creating truck: %v", err)
		return
	}
	
	fmt.Printf("✓ %s\n", truck.GetInfo())
	fmt.Printf("  %s\n", truck.Start())
	fmt.Printf("  %s\n", truck.Stop())
	fmt.Println()
}

func demonstrateErrorHandling(factory *products.VehicleFactory) {
	fmt.Println("2. Error Handling:")
	fmt.Println("------------------")

	// Try to create an unsupported vehicle type
	_, err := factory.CreateVehicle("airplane", "Boeing", "747")
	if err != nil {
		fmt.Printf("✗ Expected error: %v\n", err)
	}
	fmt.Println()
}

func demonstrateSupportedTypes(factory *products.VehicleFactory) {
	fmt.Println("3. Supported Vehicle Types:")
	fmt.Println("---------------------------")

	supportedTypes := factory.GetSupportedTypes()
	for i, vehicleType := range supportedTypes {
		fmt.Printf("%d. %s\n", i+1, vehicleType)
	}
	fmt.Println()
}

func demonstrateCustomTruck(factory *products.VehicleFactory) {
	fmt.Println("4. Custom Truck with Specific Capacity:")
	fmt.Println("---------------------------------------")

	// Create a truck with custom capacity
	heavyTruck, err := factory.CreateVehicleWithCapacity(products.TruckType, "Volvo", "FH16", 25)
	if err != nil {
		log.Printf("Error creating heavy truck: %v", err)
		return
	}
	
	fmt.Printf("✓ %s\n", heavyTruck.GetInfo())
	fmt.Printf("  %s\n", heavyTruck.Start())
	fmt.Printf("  %s\n", heavyTruck.Stop())
	fmt.Println()
}