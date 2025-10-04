package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup is used to wait for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.
var wg sync.WaitGroup

// wg := new(sync.WaitGroup) way to create wait group pointer
// var wg *sync.WaitGroup = &sync.WaitGroup{} another way to create wait group pointer

func Worker(id int) {
	defer wg.Done()
	fmt.Println("Worker", id, "started")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "done")
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(i)
	}
	wg.Wait()
	fmt.Println("Main code done")
}

// Output random but wait till last all go routine finishes it works

// without wait group it will not wait for all go routines to finish

// func Worker(id int) {

// 	fmt.Println("Worker", id, "started")
// 	time.Sleep(time.Second)
// 	fmt.Println("Worker", id, "done")
// }

// func main() {
// 	for i := 0; i < 5; i++ {

// 		go Worker(i)
// 	}

// 	fmt.Println("Main code done")
// }
