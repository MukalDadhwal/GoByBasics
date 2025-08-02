package main

import (
	"fmt"
	// "time"
)

func SelectProgram() {
	messages := make(chan string)
	signals := make(chan bool)

	// fmt.Println(<-messages)
	// fmt.Println(<-signals) // result in deadlock as both are unbuffered so blocking in nature

	/* Try uncommenting this
	go func(){
		messages<-"hello";
	}()
	
	time.Sleep(1 * time.Second)
	*/

	select{
	case msg := <-messages:
		fmt.Println("Received msg: ", msg)
	case signal := <-signals:
		fmt.Println("Received msg: ", signal)
	default:
		fmt.Println("No message received")
	}	// here select will chose the one which executes first here as both have nothing to receive the default is executed

	close(messages)
	close(signals)
}
