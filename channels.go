package main

import (
	"fmt"
	"time"
)

// Rule of thumb: An unbuffered channel needs both a sender and receiver to be ready at the same time i.e concurrently—in different
// goroutines—for anything to happen.

// In go unbuffered channels have zero size meaning they cannot store anything therefore they must after sending a msg
// immediately receive it. If not done the thread/routine is blocked and a deadlock occurs if no threads are free.

func addToChannel(msg string, channel chan string){
	channel<-msg
}

func receiveFromChannel(channel chan string){
	fmt.Println(<-channel)
}

func worker(channel chan bool){
	fmt.Println("working...")
	time.Sleep(3 * time.Second)
	fmt.Println("done")

	channel<-true
}

func ChannelProgram(){
	msgChannel := make(chan string)  // unbuffered go channel having size 0

	go addToChannel("hey", msgChannel) 
	// go receiveFromChannel(msgChannel) // till this routine started executing the main exited
	fmt.Println(<-msgChannel)

	// buffered channel
	buf := make(chan string, 2)

	buf<-"everyday"
	buf<-"is" // will not go in deadlock as it does not require a receive immediately

	fmt.Println(<-buf)
	fmt.Println(<-buf)

	fmt.Println("done")

	// example: waiting on a routine to complete
	channel := make(chan bool)

	go worker(channel)

	if <-channel {
		fmt.Println("task completed.")
	}
	 
	fmt.Println("main exiting")

}
