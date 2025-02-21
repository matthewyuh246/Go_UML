package main

import "fmt"

type Shape interface {
	getArea() int
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) setWidth(width int) {
	r.width = width
}

func (r *Rectangle) setHeight(height int) {
	r.height = height
}

func (r *Rectangle) getArea() int {
	return r.height * r.width
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

type Square struct {
	length int
}

func (s *Square) setLength(length int) {
	s.length = length
}

func (s *Square) getArea() int {
	return s.length * s.length
}

func NewSquare(length int) *Square {
	return &Square{
		length: length,
	}
}

func f(shape Shape) {
	fmt.Println(shape.getArea())
}

// func main() {
// 	r1 := NewRectangle(3, 4)
// 	f(r1)

// 	r1.setWidth(4)
// 	r1.setHeight(5)
// 	f(r1)

// 	r2 := NewSquare(5)
// 	f(r2)
// }
