package main

import (
	"context"
	"fmt"
	"time"
)

// In a go programming lang context is a pkg that provides a way to carry deadlines a, cancelation and signals across the API
// boundary and b/w the processes.
// It is used to manage the lifecycle of a req and to cancel the req if it takes too long to execute or if the client cancels the req.

// dummy function to demonstrate the use of context

func WorkSomething(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Cancelled", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // cancel the context and release the resouces when done

	// fire a goroutine to perform some processing
	go WorkSomething(ctx)

	//wait for few sec to do the complete lifecycle of work
	time.Sleep(3 * time.Second) //output: Cancelled context deadline exceeded

}
