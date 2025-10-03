package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex is used to provide a locking mechanism to ensure that only one goroutine can access a critical section of code at a time.
// Analogy let us suppose a person locks his bathroom door while taking a bath so that no one else can enter until he is done.
// In Go, we can use sync.Mutex to achieve this.

// However, in this example, we are not using mutex, which may lead to race conditions.
// output may vary on different runs.
// var counter int

// func increment() {
// 	for i := 0; i < 1000; i++ {
// 		counter++ // Critical section multiple goroutines may access this simultaneously
// 	}
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go increment()

// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println("Final Counter:", counter)
// }

var counter int
var mu sync.Mutex

func increment() {

	for i := 0; i < 1000; i++ {
		mu.Lock() // Lock the mutex before entering the critical section
		counter++
		mu.Unlock() // Unlock the mutex after leaving the critical section
	}
}
func main() {
	for i := 0; i < 5; i++ {
		go increment()
	}
	time.Sleep(time.Second)
	fmt.Println("Final Counter:", counter)
}
