package main

import (
	"fmt"
)

func counter() func() int{
	count := 0;
	return func() int {
		count++;
		return count;
	}
}

func ClosureProgram(){

	nextInt := counter();

	nextInt();
	nextInt();
	fmt.Println(nextInt()); // 3

	newNextInt := counter(); // new value of counter is intialised with value 0

	newNextInt();
	fmt.Println(newNextInt());
	
}