package main

import (
	"fmt"
)

// a structure is a collection of fields
type person struct {
	name string
	age  int
}

// method by reference
func (p *person) changeAge(newAge int) {
	p.age = newAge
}

// method by value
func (p person) changeAge3(newAge int) person{
	p.age = newAge;
	return p;
}

// Methods can be defined for either pointer or value receiver types
func (p person) changeAge2(newAge int) *person {
	p.age = newAge
	return &p
}

func StructProgram() {
	p1 := person{name: "Rajan", age: 59}
	fmt.Println(p1)

	p2 := person{"Hayat", 40}
	fmt.Println(p2)

	// anonmous struct (used for normal grouping of data)
	dog := struct {
		name, color string
	}{
		"Boo",
		"White",
	}

	fmt.Println(dog)

	p1.changeAge(20)
	fmt.Println(p1)

	p2.changeAge(30)
	fmt.Println(p2)

	p1 = *p1.changeAge2(24)
	fmt.Println(p1)

	p1 = p1.changeAge3(100)
	fmt.Println(p1)

}
