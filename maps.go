package main

import (
	"fmt"
	"maps"
)

func MapsProgram() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"] // if value not present returns the default for int i.e 0
	fmt.Println("v3:", v3)

	fmt.Println("len: ", len(m))

	// remvoing keys
	delete(m, "k2")

	clear(m)

	_, prs := m["k2"] // right way of accessing values from map _ = value prs = boolean
	fmt.Println("prs:", prs)

	// defining maps in one line
	n := map[string]int{"foo": 1, "bar": 2}

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
