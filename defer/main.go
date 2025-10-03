package main

import "fmt"

func main() {
	defer fmt.Println("Ending main function")              // Now this will be called at the end of main function
	fmt.Println("Starting main function")                  // synchronous print
	defer fmt.Println("Starting main function with defer") // This will be called before the previous defer due to LIFO order

	// loop to print numbers from 0 to 4
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%d\n", i)
	// }

	// defer keyword in loop

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // This will print 4 3 2 1 0 due to LIFO order
	}
	fmt.Println("Loop ended") // synchronous print
}

// Defer keyword is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

// When a function returns, its deferred calls are executed in last-in-first-out order. (LIFO	)

//Starting main function
// Loop ended
// 4 3 2 1 0 Starting main function with defer
// Ending main function
