package main

import "fmt"

func main() {
	fmt.Println("Learning Go arrays")
	numbers := [10]int{2, 5, 8, 17, 64, 1, 23} // array of fixed size 10
	fmt.Println("Array of number is:", numbers)
	fmt.Println("Length of array is:", len(numbers))
	fmt.Println("Indexing of 1st element is:", numbers[0])
	fmt.Println("Lenght of array:", len(numbers))

	// iterating through array
	for i := 0; i <= len(numbers)-1; i++ {
		fmt.Println(numbers[i])
	}

	// range based for loop
	for index, value := range numbers {
		fmt.Println("Index is", index, "Value is", value)
	}

	// for i (index) , val := range iterable{
	// 	// do something
	// }

	// what range do ?
	// it returns two values, index and value of that index
	// if you want to ignore any of the two values, use _ (blank identifier)
	for _, val := range numbers {
		fmt.Println("Value is", val)
	}
	// what %d %s %T do?
	// %d is for integers
	// %s is for strings
	// %T is for type of variable
	for _, val := range numbers {
		fmt.Printf("Value is %d and Type is %T\n", val, val)
	}

}
