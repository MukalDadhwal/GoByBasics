package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 5
}

// variable length args
func sum(nums ...int) int {
	total := 0

	for i := range len(nums) {
		total += nums[i]
	}

	return total
}

func FunctionProgram() {

	result := plus(2, 3)
	fmt.Println(result)

	a, b := vals() // we don't want a value we can use _
	fmt.Println(a, b)

	nums := []int{3, 4, 5, 6}
	fmt.Println(sum(nums...))

	// 1. ANONYMOUS FUNCTIONS
	hello := func() {
		fmt.Println("Hello there!")
	}

	// 2.
	func(msg string) {
		fmt.Println(msg)
	}("HELLO GO")

	hello()

	// Anonymous recursive functions
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(10))

}
