package main

import (
	"fmt"
	"math"
)

type circle struct{
	radius float64;
}

// circumference 
func (c circle) circumference() float64{
	return 2 * 3.14 * float64(c.radius);
}

// area
func (c circle) area() float64{
	return math.Pi * math.Pow(float64(c.radius), 2);
}

func StructCircleProgram(){
	var radius int

	fmt.Print("Enter the radius: ")
	fmt.Scan(&radius)

	c := circle{radius: float64(radius)}

	// Sprintf -> returns the formatted string
	// Printf -> prints the formatted string 

	fmt.Printf("Circumference: %.2f\n", c.circumference())
	fmt.Printf("Area: %.2f", c.area())
}
