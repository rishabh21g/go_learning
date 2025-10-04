package main

import "fmt"

func main() {
	chan1 := make(chan int) // Create a channel of type int
	//fmt.Println(<-chan1)
	// fmt.Println(<-chan1)
	//Here is a deadlock because the main goroutine is trying to send a value into the channel but there is no other goroutine to receive it.
	// Fixing deadlock using goroutine
	go func() {
		chan1 <- 1342
	}()

	fmt.Println(<-chan1)
}

/*
___________________________________________________________________________________________________________
A channel is a pipe that allows goroutines to communicate with each other.
You can send values into a channel and receive values from it.
Channels are typed → a chan int can only carry int, a chan string only string, etc.
Think of it as a conveyor belt in a factory:
Workers (goroutines) put things on it.
Other workers pick them up.
_______________________________________________________________________________________________________
Channel Directions
By default, channels are bidirectional.
But you can restrict them:
chan<- → send-only
<-chan → receive-only
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.
Real-World Use Cases
Worker pools (distributing tasks to goroutines).
Fan-out / Fan-in (multiple senders → one receiver, or one sender → multiple receivers).
Pipelines (processing data in stages).
Cancellation signals (close a channel to tell goroutines to stop).
Time-based operations (time.After, time.Tick use channels).
*/
