package main

import "fmt"

type Vehicle interface {
	GetNumberOfWheels() int32
	GetSeatingCapcity() int32
}

func GetVehicleObject(vehicleType string) Vehicle {
	switch vehicleType {
	case "Car":
		return &Car{}
	case "Bike":
		return &Bike{}
	}
	return nil
}

type Car struct {
	numberOfWheels int32
	seatingCapcity int32
}

func (t *Car) GetNumberOfWheels() int32 {
	return int32(4)
}

func (t *Car) GetSeatingCapcity() int32 {
	return int32(5)
}

type Bike struct {
	numberOfWheels int32
	seatingCapcity int32
}

func (t *Bike) GetNumberOfWheels() int32 {
	return int32(2)
}

func (t *Bike) GetSeatingCapcity() int32 {
	return int32(2)
}

func main() {
	c := GetVehicleObject("Car")
	fmt.Println("No of wheels in car :", c.GetNumberOfWheels())
	fmt.Println("No of max people in car :", c.GetSeatingCapcity())

	b := GetVehicleObject("Bike")
	fmt.Println("No of wheels in bike :", b.GetNumberOfWheels())
	fmt.Println("No of max people on bike :", b.GetSeatingCapcity())
}
