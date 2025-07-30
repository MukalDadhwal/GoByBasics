package main

import (
	// "fmt"
	"time"
)

func Switch() string {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		return "It's before noon"

	default:
		return "It's after noon"
	}
}
