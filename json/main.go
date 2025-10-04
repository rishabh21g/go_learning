package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	StudentID int    `json:"student_id"`
	FullName  string `json:"full_name"`
	Age       int    `json:"age"`
	IsActive  bool   `json:"is_active"`
}

func main() {
	fmt.Println("Learning JSON in Golang")

	//  Marshaling (Go struct → JSON)
	std1 := Student{
		StudentID: 1,
		FullName:  "Rishabh Gupta",
		Age:       23,
		IsActive:  true,
	}

	encodedJSON, err := json.MarshalIndent(std1, "", "  ") // pretty print
	if err != nil {
		fmt.Printf("Error while marshaling: %v\n", err)
		return
	}
	fmt.Println("Encoded JSON:")
	fmt.Println(string(encodedJSON))

	//  Unmarshaling (JSON → Go struct)
	var std2 Student
	jsonStr := `{"student_id":1,"full_name":"Rishabh Gupta","age":23,"is_active":true}`

	if err := json.Unmarshal([]byte(jsonStr), &std2); err != nil {
		fmt.Printf("Error while unmarshaling: %v\n", err)
		return
	}

	fmt.Println("Decoded Struct:")
	fmt.Printf("%+v\n", std2) // %+v shows field names too

	// When we does'nt know exact field we often use map and interfaces
	// 	map = key-value pair hota hai.
	// string = JSON me jo keys hoti hain (e.g. "id", "name", "active") wo string hoti hain.
	// interface{} = iska matlab hai kuch bhi type (int, float, string, bool…)

	//Unmarshal Using map and interfaces

	var result map[string]interface{}
	jsonString := `{"id":101,"active":true,"name":"Laptop"}`
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		fmt.Printf("Error while unmarshaling: %v\n", err)
		return
	}
	fmt.Printf("Dynamic struct using map interfaces: %v\n", result)

}

/* JSON - JavaScript Object Notation Lightweight format for exchanging data (text-based, human-readable). Used for client server interaction for data transfer Go provides encoding/json package for marshaling and unmarshaling Marshaling - Converting Go objects to JSON format Unmarshaling - Converting JSON data to Go objects */
