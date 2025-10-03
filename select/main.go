package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning Select in Go")

	// Creating a channel
	chan1 := make(chan string)
	chan2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "Message for channel 1"
	}()
	// Goroutine 2
	go func() {
		time.Sleep(1 * time.Second)
		chan2 <- "Message for channel 2"
	}()

	//select statement is used to wait on multiple channel operations.
	// It blocks until one of its cases can run, then it executes that case.
	// If multiple cases are ready, one case is chosen at random to execute.
	// If none are ready, it blocks until one becomes ready.
	// A default case, if present, is run if no other case is ready.
	// The select statement is similar to a switch statement, but for channels.
	// Each case in a select statement must be a channel operation (send or receive).
	// The select statement will block until one of the channel operations is ready to proceed.
	select {
	case msg1 := <-chan1:
		fmt.Println(msg1)
	case msg2 := <-chan2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: no message received")
	}

	// output: Randomly either "Message for channel 1" or "Message for channel 2" or "Timeout: no message received" if neither message is received within 3 seconds if withou no sleep in goroutines

	// output: "Message for channel 1" if with 1 second sleep in goroutine 1 and no sleep in goroutine 2
	// output: "Message for channel 2" if with 1 second sleep in goroutine 2 and no sleep in goroutine 1
}
