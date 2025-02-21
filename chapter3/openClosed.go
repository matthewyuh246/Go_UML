package main

import (
	"math"
)

type IEmployee interface {
	GetBonus(base int) int
}

type JuniorEmployee struct {
	name string
}

func NewJuniorEmployee(name string) IEmployee {
	return &JuniorEmployee{
		name: name,
	}
}

func (J *JuniorEmployee) GetBonus(base int) int {
	return int(math.Floor(float64(base) * 1.1))
}

// func main() {
// 	emp1 := NewJuniorEmployee("Yamada")
// 	base := 100
// 	fmt.Println(emp1.GetBonus(base))
// }
