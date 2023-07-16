package main

import (
	"fmt"
	"time"
)

func main() {
	// result := fetchResource()
	// fmt.Println("the result: ", result)
	//* 1) Unbuffered channel
	// Using an unbuffered channel will throw errors in this case as it has no space to store anything
	//cookie analogy, unbuffered chan can't store any cookies
	//! A channel in golang will always block if it is full
	//! The moment we write something to an unbuffered channel, it is FULL and it BLOCKS
	// resultch := make(chan string)
	// resultch <- "foo"
	//Code below will not execute at all

	/*
	* SO how to use unbuffered channels and still get the job done?
	 */
	resultch := make(chan string)

	//* Here we've created an empty hand to accept a cookie, since we cannot store a cookie in an unbuffered channel
	//* we need someone to immediately pass it to.
	//* This goroutine will run in a separate thread, so it is ready to accept.
	go func() {

		result := <-resultch
		fmt.Println(result)

	}()
	//* Once we have a receiver ready, we can send the cookie in the channel and it will work!

	resultch <- "foo"

	// *2) Buffered channel
	// *Buffered channel can store 10 cookies in this case / 10 strings
	buffchan := make(chan string, 10)

	buffchan <- "cookie"
	resultBuff := <-buffchan
	fmt.Println(resultBuff)

	// go fetchResource(1) //async call

	//Anonymous call, same thing as above
	// go func() {
	// 	fetchResource(1)
	// }()
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
