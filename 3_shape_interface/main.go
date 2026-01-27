package main

import (
	"fmt"
	"math"
)

// Shape interface defines the Area method
type Shape interface {
	Area() float64
}

// Rectangle struct with Width and Height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct with Radius
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// PrintArea accepts any Shape and prints its area
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	fmt.Println("=== Shape Interface Example ===")

	// Create a rectangle
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle (Width: %.1f, Height: %.1f)\n", rect.Width, rect.Height)
	PrintArea(rect)

	fmt.Println()

	// Create a circle
	circle := Circle{Radius: 7}
	fmt.Printf("Circle (Radius: %.1f)\n", circle.Radius)
	PrintArea(circle)

	fmt.Println()

	// Demonstrate polymorphism
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Rectangle{Width: 8, Height: 2},
	}

	fmt.Println("All shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d - ", i+1)
		PrintArea(shape)
	}
}
