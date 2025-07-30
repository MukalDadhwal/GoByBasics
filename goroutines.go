package main

import (
	"fmt"
	"time"
)

func foo(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func GoRoutineProgram() {
	foo("direct")

	go foo("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("anon")

	time.Sleep(time.Second)
    fmt.Println("done")

}
