package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// type User struct {
// 	username    string
// 	email       string
// 	age         int
// 	active      bool
// 	skills      []string
// 	collegeName string
// 	favLanguage string
// }

type User struct {
	ID       string
	Name     string
	Email    string
	Age      int
	password string
}

func randomPasswordGenerator(passLength int) string {
	const passwordCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	var generatedPassword string
	for i := 0; i < passLength; i++ {
		generatedPassword = generatedPassword + string(passwordCharset[rand.Intn(len(passwordCharset))])

	}
	return generatedPassword
}

func main() {
	fmt.Println("Learning Go structs")
	id := rand.Int() * 1000000
	password := randomPasswordGenerator(15)

	user1 := User{}
	user1.Name = "Sanchay Roy"
	user1.Email = "sanchayroy@gmail.com"
	user1.ID = strconv.FormatInt(int64(id), 10)
	user1.Age = 22
	user1.password = password

	// fmt.Printf("User struct defined: %+v\n", user1)
	var u User
	u2 := u
	u2.ID = "1" // makes a copy (value semantics) does not affect u
	u2.Name = "Alice"
	u2.Email = "abc@gmail.com"
	// fmt.Printf("Empty User struct: %#v\n", u) // %#v will print the field names as well as the struct name
	// fmt.Printf("Empty User struct copy: %#v\n", u2)

	// difference in u and u2?
	// on the basis of memory address
	fmt.Printf("Memory address of u: %p\n", &u) // & gives the memory address of the variable %p is used to print the memory address %d is used to print the integer value %s is used to print the string value
	fmt.Printf("Memory address of u2: %p\n", &u2)

}

// What are structs?
// Structs are collections of fields
// They are used to group data together to form records
// They can be used to create complex data structures
// user1 := User{
// 	username:    "Rishabh Gupta",
// 	email:       "biasedengineer@gmail.com",
// 	age:         23,
// 	active:      true,
// 	skills:      []string{"Golang", "Javascript", "Python", "NOdeJS", "ReactJS", "NextJS"},
// 	collegeName: "IIT Madras",
// 	favLanguage: "Golang",
// }
// fmt.Println("User1 Details:", user1)
//User1 Details: {Rishabh Gupta biasedengineer@gmail.com 23 true [Golang Javascript Python NOdeJS ReactJS NextJS HTML CSS MongoDB SQL] IIT Madras Golang}
// fmt.Printf("User1 Details: %+v\n", user1 // %+v will print the field names as well
// fmt.Printf("User1 Details: %#v\n", user1) )// %#v will print the field names as well as the struct name
// A struct value is a contiguous block of memory holding its fields (or references to other data like slices/maps/pointers).

// In structs no pvt public concept for importing/exporting packages
// Only capitalized fields are exported
// if field starts with small letter then it is unexported
