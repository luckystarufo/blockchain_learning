package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	fmt.Printf("Rectangle width: %.2f, height: %.2f\t", r.Width, r.Height)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("Circle radius: %.2f\t", c.Radius)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
}
