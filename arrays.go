package main

import "fmt"

func ArrayProgram() {

	// declare array
	var arr [10]int

	fmt.Println(arr)

	// updating
	arr[0] = 100
	arr[1] = 101
	// ...

	value := 101

	for i := range 10 {
		arr[i] = value
		value++
	}

	// accessing
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	// length
	fmt.Println(len(arr))

	arr = [10]int{1, 2, 3, 4, 5} // can redeclare with = but size should be same

	fmt.Println(arr)

	arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // if size is not given (compiler will infer) the array must of size 10

	fmt.Println(arr)

	arr2 := [...]int{1,2,10:400,500}

	fmt.Println(arr2)

}
