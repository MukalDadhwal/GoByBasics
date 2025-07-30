package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length  float64
	breadth float64
}

type circle2 struct {
	radius float64
}

// defining area and perimeter on rect
func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// defining area and perimeter on circle
func (c circle2) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle2) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// multi utility function
func someFunc(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func findType(g geometry){
	switch g.(type){
		case rect:
			fmt.Println("Type is rect")
		case circle2:
			fmt.Println("Type is circle")
		default:
			fmt.Println("Type is not rect or circle")
	}
}

func InterfaceProgram() {
	// both structs implements all the functions of geometry interface which automatically implements the interface
	r := rect{10, 12}
	c := circle2{5}

	someFunc(r)
	someFunc(c)

	findType(r)
	findType(c)
}
