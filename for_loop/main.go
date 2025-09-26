package main

import "fmt"

func main() {
	fmt.Println("Learning Go loops")

	// for loop is the only loop in Go
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// while loop with for syntax
	var num int = 1
	for num < 10 {
		num += num
		fmt.Println("Value of num is:", num)
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")

	// }

	// for loop with break and continue keywords
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
