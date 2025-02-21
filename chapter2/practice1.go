package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary int
}

func (e *Employee) Work() {
	fmt.Println("働きます")
}

func (e *Employee) getSalary() int {
	return e.salary
}

func (e *Employee) setSalary(salary int) {
	e.salary = salary
}
