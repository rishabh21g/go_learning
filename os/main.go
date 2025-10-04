package main

import (
	"fmt"
	"os"
)

func main() {
	// Read write and create package
	file, err := os.Create("user.txt")
	if err != nil {
		fmt.Println("Error while creating file", err.Error())
	}
	defer file.Close()
	file.WriteString("This is my first file using OS package")
	fmt.Println("File created succesfully!")
}
