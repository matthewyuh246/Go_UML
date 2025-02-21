package main

import "fmt"

type Movable interface {
	start()
	stop()
}

type Flyable interface {
	fly()
}

type Vehicle struct {
	name  string
	color string
}

type Airplane struct {
	Vehicle
}

func NewAirplane(name, color string) *Airplane {
	return &Airplane{
		Vehicle: Vehicle{
			name:  name,
			color: color,
		},
	}
}

func (a *Airplane) start() {
	fmt.Println("start!")
}

func (a *Airplane) stop() {
	fmt.Println("stop!")
}

func (a *Airplane) fly() {
	fmt.Println("fly!")
}

type Car struct {
	Vehicle
}

func NewCar(name, color string) *Car {
	return &Car{
		Vehicle: Vehicle{
			name:  name,
			color: color,
		},
	}
}

func (c *Car) start() {
	fmt.Println("start!")
}

func (c *Car) stop() {
	fmt.Println("stop!")
}

// func main() {
// 	v1 := NewAirplane("AirBus", "white")
// 	v2 := NewCar("Prius", "black")

// 	v1.fly()
// 	v2.start()
// }
