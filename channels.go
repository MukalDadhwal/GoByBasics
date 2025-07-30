// channels in go
package main

import "fmt"

// Rule of thumb: An unbuffered channel needs both a sender and receiver to be ready at the same time i.e concurrently—in different 
// goroutines—for anything to happen. In simple words unbuffererd channels needs to be synchronised

func addToChannel(msg string, channel chan string){
	channel<-msg
}

func receiveFromChannel(channel chan string){
	fmt.Println(<-channel)
}

func ChannelProgram(){
	msgChannel := make(chan string)  // unbuffered go channel

	go addToChannel("hey", msgChannel)
	// go receiveFromChannel(msgChannel)
	fmt.Println(<-msgChannel)
	
	var inp string
	fmt.Scan(&inp)	
	fmt.Println("done")
}