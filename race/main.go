package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// what issue with this code?
// temp variable is shared between all goroutines, so they are all
//
//	modifying the same variable, which can lead to a race condition. This can cause the final
//
// value of counter to be less than num, which is not the expected result. To fix this issue, we can
// use a mutex to synchronize access to the counter variable, or we can use atomic operations to
//
//	increment the counter safely without the need for a mutex.
// func main() {
// 	var wg sync.WaitGroup
// 	const num = 100
// 	counter := 0
// 	wg.Add(num)
// 	for i := 0; i < num; i++ {
// 		go func() {
// 			defer wg.Done()
// 			temp := counter
// 			runtime.Gosched() // Yield the processor to allow other
// 			// goroutines to run and increase the chance of a race condition
// 			temp++
// 			counter = temp
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Code Executed: ", counter)

// }

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	const num = 100
	counter := 0
	wg.Add(num)
	time.Now()
	for range num {
		go func() {
			defer wg.Done()
			mu.Lock()
			temp := counter
			runtime.Gosched() // Yield the processor to allow other
			// goroutines to run and increase the chance of a race condition incrementing the counter without the mutex
			temp++
			counter = temp
			mu.Unlock()
		}()
	}
	wg.Wait()
	duration := time.Since(time.Now())
	fmt.Println("Code Executed: ", counter)
	fmt.Println("Duration: ", duration)
}
