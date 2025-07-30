package main

import "fmt"

func Conditionals(grade int) {

	if grade >= 85 && grade <= 100 {
		fmt.Println("Your grade is A")
	} else if grade >= 75 && grade < 85 {
		fmt.Println("Your grade is B")
	} else {
		fmt.Println("Your grade is ASS")
	}

}
