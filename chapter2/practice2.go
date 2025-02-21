package main

import "math"

type Shape interface {
	calcArea() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type IRectangle interface {
	calcArea() float64
}

func NewRectangle(width, height float64) IRectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r *Rectangle) calcArea() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

type ICircle interface {
	calcArea() float64
}

func NewCircle(radius float64) ICircle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) calcArea() float64 {
	return c.radius * c.radius * math.Pi
}
