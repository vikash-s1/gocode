package products

import (
	"errors"
	"strings"
)

// VehicleType represents the type of vehicle to create
type VehicleType string

const (
	CarType        VehicleType = "car"
	MotorcycleType VehicleType = "motorcycle"
	TruckType      VehicleType = "truck"
)

// VehicleFactory is the factory struct that creates vehicles
type VehicleFactory struct{}

// NewVehicleFactory creates a new instance of VehicleFactory
func NewVehicleFactory() *VehicleFactory {
	return &VehicleFactory{}
}

// CreateVehicle is the factory method that creates vehicles based on type
func (vf *VehicleFactory) CreateVehicle(vehicleType VehicleType, brand, model string) (Vehicle, error) {
	switch strings.ToLower(string(vehicleType)) {
	case string(CarType):
		return &Car{
			Brand: brand,
			Model: model,
		}, nil
	case string(MotorcycleType):
		return &Motorcycle{
			Brand: brand,
			Model: model,
		}, nil
	case string(TruckType):
		return &Truck{
			Brand:    brand,
			Model:    model,
			Capacity: 10, // Default capacity
		}, nil
	default:
		return nil, errors.New("unknown vehicle type: " + string(vehicleType))
	}
}

// CreateVehicleWithCapacity creates a truck with specific capacity
func (vf *VehicleFactory) CreateVehicleWithCapacity(vehicleType VehicleType, brand, model string, capacity int) (Vehicle, error) {
	if vehicleType != TruckType {
		return vf.CreateVehicle(vehicleType, brand, model)
	}
	
	return &Truck{
		Brand:    brand,
		Model:    model,
		Capacity: capacity,
	}, nil
}

// GetSupportedTypes returns all supported vehicle types
func (vf *VehicleFactory) GetSupportedTypes() []VehicleType {
	return []VehicleType{CarType, MotorcycleType, TruckType}
}