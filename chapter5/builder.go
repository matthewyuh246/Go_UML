package main

import "fmt"

type Computer struct {
	Type string
	cpu  string
	ram  int
}

type ComputerBuilder interface {
	AddCpu(cpu string)
	AddRam(ram int)
}

type DesktopBuilder struct {
	computer *Computer
}

func NewDesktopBuilder() *DesktopBuilder {
	return &DesktopBuilder{
		&Computer{
			Type: "Desktop",
		},
	}
}

func (db *DesktopBuilder) AddCpu(cpu string) {
	db.computer.cpu = cpu
}

func (db *DesktopBuilder) AddRam(ram int) {
	db.computer.ram = ram
}

func (db *DesktopBuilder) getResult() Computer {
	return *db.computer
}

type LaptopBuilder struct {
	computer *Computer
}

func NewLaptopBuilder() *LaptopBuilder {
	return &LaptopBuilder{
		&Computer{
			Type: "Laptop",
		},
	}
}

func (db *LaptopBuilder) AddCpu(cpu string) {
	db.computer.cpu = cpu
}

func (db *LaptopBuilder) AddRam(ram int) {
	db.computer.ram = ram
}

func (db *LaptopBuilder) getResult() Computer {
	return *db.computer
}

type Director struct {
	builder ComputerBuilder
}

func NewDirector(builder ComputerBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) construct() {
	d.builder.AddCpu("Core i5")
	d.builder.AddRam(16)
}

func (d *Director) highSpecConstruct() {
	d.builder.AddCpu("Core i9")
	d.builder.AddRam(64)
}

func main() {
	desktopBuilder := NewDesktopBuilder()
	desktopDirector := NewDirector(desktopBuilder)
	desktopDirector.construct()
	desktopComputer := desktopBuilder.getResult()
	fmt.Println(desktopComputer)

	laptopBuilder := NewLaptopBuilder()
	laptopDirector := NewDirector(laptopBuilder)
	laptopDirector.highSpecConstruct()
	laptopComputer := laptopBuilder.getResult()
	fmt.Println(laptopComputer)
}
