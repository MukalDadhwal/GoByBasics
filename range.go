package main

import (
	"fmt"
)

func RangeProgram() {

	nums := []int{1, 2, 3, 4, 5}

	for i := range len(nums) {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// _ = index j = element
	for _, j := range nums { // the first value given by range is the index but on inbuilt types like map, slices, arrays it will return the element as well
		fmt.Print(j, " ")
	}
	fmt.Println()

	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs { // here the first element is key and sec is value
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
