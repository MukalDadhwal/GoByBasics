package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) (bool, error) {
	if age < 0 {
		return false, errors.New("INVALID AGE PASSED")
	}

	if age >= 18 {
		return true, nil
	} else {
		return false, nil
	}
}

func ErrorProgram() {
	r, e := checkAge(30)

	fmt.Println("Result:", r, " Error:", e)

	r1, e1 := checkAge(-30)

	fmt.Println("Result:", r1, " Error:", e1)
}
