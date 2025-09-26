package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	/*
		A pointer is a variable that stores the memory address of another variable.
		In Go, pointers are represented using the asterisk (*) symbol.

		Here's a simple example to illustrate the concept of pointers in Go:
	*/

	var x int = 42
	var ptr *int = &x // ptr is a pointer to an integer, and it holds the address of x

	fmt.Println("Value of x:", x)                 // Output: Value of x: 42
	fmt.Println("Address of x:", &x)              // Output: Address of x:
	fmt.Println("Value of ptr:", ptr)             // Output: Value of ptr:
	fmt.Println("Value pointed to by ptr:", *ptr) // Output: Value pointed to by ptr: 42
}
