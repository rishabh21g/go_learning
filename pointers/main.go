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

	//  pointer_name * Data_type Operator
	// & is the address operator, which returns the memory address of a variable.
	// * is the dereference operator, which returns the value stored at the memory address that a pointer points to.

	// Pointers are useful for various reasons, such as:
	// 1. They allow you to modify the value of a variable from within a function.
	// 2. They enable you to work with large data structures without copying them.
	// 3. They can be used to create complex data structures like linked lists and trees.

	// It's important to note that Go has garbage collection, so you don't need to worry about memory management when
	// using pointers. However, you should still be cautious when working with pointers to avoid issues like nil pointer
	// dereferences or memory leaks.
}
