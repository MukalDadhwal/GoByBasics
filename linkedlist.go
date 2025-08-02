package main

import (
	"fmt"
)

type Node[T any] struct{
	next *Node[T]
	val T
}

func LinkedList(){
	n1 := Node[int]{val: 10}
	n2 := Node[int]{val: 20}
	n3 := Node[int]{val: 30}

	n1.next = &n2;
	n2.next = &n3;

	// Traverse and print values
    current := &n1
    for current != nil {
        fmt.Println(current.val)
        current = current.next
    }

}	