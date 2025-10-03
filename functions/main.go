package main

import "fmt"

func increment(x int) int {
	return x + 1
}
func main() {
	fmt.Println("Learning Go functions!")
	// A function is a piece of code that performs a specific task.
	// It can take inputs, process them, and return an output.
	// Functions help in organizing code, making it reusable and easier to read.

	// Example usage of the increment function
	x := 20
	result := increment(x)
	fmt.Println("Incremented value:", result)
}
