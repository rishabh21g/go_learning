package main

import "fmt"

func main() {
	fmt.Println("Learning slices in go")
	/* Slice is a dynamically sized, flexible view into the elements of an array.
	   Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	   To create an empty slice with non-zero length and capacity, use the built-in make function:
		`make([]T, length, capacity)`
	   where T is the element type of the slice.

	   An uninitialized slice is nil and has a length and capacity of 0:
		var s []int
		fmt.Println(s, len(s), cap(s)) // [] 0 0 */
	// emptySlice := []int{}
	// fmt.Println("Empty Slice:", emptySlice)
	// fmt.Println("Length of Empty Slice:", len(emptySlice))
	// fmt.Println("Capacity of Empty Slice:", cap(emptySlice))

	// Creating a slice using make function
	madeSlice := make([]string, 5, 10) // T is the element type of the slice
	fmt.Println("Made Slice:", madeSlice)
	fmt.Println("Length of Made Slice:", len(madeSlice))
	fmt.Println("Capacity of Made Slice:", cap(madeSlice))

	/* Difference b/w len and capacity
	The length of a slice is the number of elements it contains.
	 The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	 The capacity is always greater than or equal to the length.*/

	arraSlice := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := arraSlice[1:4] // creates a slice from index 1 to 3 (4 is excluded)
	fmt.Println("Slice from Array:", sliceFromArray)
	fmt.Println("Length of Slice from Array:", len(sliceFromArray))
	fmt.Println("Capacity of Slice from Array:", cap(sliceFromArray)) // capacity is from index 1 to end of array

}
