package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

type Motorcycle struct{}

func (m *Motorcycle) Drive() string {
	return "Driving a motorcycle"
}

type VehicleFacoory interface {
	CreateVehicle() Vehicle
}

type CarFactory struct{}

func (cf *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

type MotorcycleFactory struct{}

func (mf *MotorcycleFactory) CreateVehicle() Vehicle {
	return &Motorcycle{}
}

type VehicleType string

const (
	CarType        VehicleType = "car"
	MotorcycleType VehicleType = "motorcycle"
)

var vehicleFactories = map[VehicleType]func() Vehicle{
	CarType:        func() Vehicle { return &Car{} },
	MotorcycleType: func() Vehicle { return &Motorcycle{} },
}

func GetVehicle(vehicleType VehicleType) Vehicle {
	if factory, ok := vehicleFactories[vehicleType]; ok {
		return factory()
	}
	return nil
}

func main() {

	car1 := GetVehicle(CarType)
	fmt.Println(car1.Drive())
	motorcycle1 := GetVehicle(MotorcycleType)
	fmt.Println(motorcycle1.Drive())

	fmt.Println("------------------------")

	carFactory := &CarFactory{}
	car := carFactory.CreateVehicle()
	fmt.Println(car.Drive())

	motorFactory := &MotorcycleFactory{}
	motorcycle := motorFactory.CreateVehicle()
	fmt.Println(motorcycle.Drive())
}
