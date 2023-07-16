package main

import "fmt"

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	msgch <- "D"

	close(msgch)

	// msg := <-msgch
	// fmt.Println("the msg is: ", msg)

	// msg = <-msgch
	// fmt.Println("the msg is: ", msg)

	//* Ranging over channels
	//* This is the consumer
	//! How will it know that we've stopped producing messages?
	// * We close the channel after we've sent all the messages
	for msg := range msgch {
		fmt.Println("the msg is: ", msg)
	}

	//* Another way to range over a channel
	msgch = make(chan string, 128)
	msgch <- "E"
	msgch <- "F"
	msgch <- "G"
	msgch <- "H"
	msgch <- "I"
	close(msgch)

	for {
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println("the msg is: ", msg)
	}

}
