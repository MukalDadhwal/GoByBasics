package main

import (
	"fmt"
	"slices"
)

func SliceProgram() {

	var s []string // unintialized arrlice (no arrize provided)

	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 10) // by default go declares the capacity to equal to length cap(s) = 10

	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// ----------------
	// small example to understand len and cap in go

	nums := make([]int, 3, 5)
	fmt.Println("Length:", len(nums))   // 3
	fmt.Println("Capacity:", cap(nums)) // 5

	nums = append(nums, 10)
	nums = append(nums, 20)
	fmt.Println("After append:", nums)  // [0 0 0 10 20]
	fmt.Println("Length:", len(nums))   // 5
	fmt.Println("Capacity:", cap(nums)) // 5

	nums = append(nums, 30)
	fmt.Println("Length:", len(nums))   // 6
	fmt.Println("Capacity:", cap(nums)) // 10 (probably)

	// ⚠️ When you exceed capacity, Go automatically allocates a new, bigger array, usually doubling the previous capacity.

	// -----------------

	// updating elements
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// append operation
	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s, "len", len(s))

	// creating new slice of same size

	c := make([]string, len(s))

	fmt.Println(c, "length: ", len(c), "capacity: ", cap(c))

	copy(c, s)

	fmt.Println(c, "length: ", len(c), "capacity: ", cap(c))

	// appending in c
	c = append(c, "g")

	fmt.Println(c, "length: ", len(c), "capacity: ", cap(c))

	// actual "slicing"

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]  		// excluding 5
	fmt.Println("sl2:", l)
 
	l = s[2:]		// till the end 
	fmt.Println("sl3:", l)

	// declaring slices in a single line
	s1 := []string {"a", "b", "c"}
	s2 := []string {"e", "f", "g"}

	if slices.Equal(s1, s2){
		fmt.Print("s1 and s2 are equal")
	}

}
