// Package basics demonstrates fundamental Go syntax and concepts
// This package covers variable declarations, data types, and basic operations
package basics

import "fmt"

// VariableExamples demonstrates different ways to declare and initialize variables in Go
func VariableExamples() {
	fmt.Println("=== Variable Declaration Examples ===")

	// Method 1: Explicit type declaration with initialization
	var name string = "Go Programming"
	var age int = 15 // Go was first released in 2009, making it ~15 years old

	// Method 2: Type inference with var keyword
	var language = "Go" // Go compiler infers this as string
	var version = 1.21  // Go compiler infers this as float64

	// Method 3: Short variable declaration (most common in functions)
	creator := "Google" // string
	isCompiled := true  // bool
	yearCreated := 2009 // int

	// Method 4: Multiple variable declarations
	var (
		projectName  = "go_learning"
		difficulty   = "Beginner"
		isOpenSource = true
	)

	// Method 5: Multiple assignments in one line
	x, y, z := 10, 20, 30

	// Printing all variables with explanatory comments
	fmt.Printf("Language: %s, Created by: %s in %d\n", name, creator, yearCreated)
	fmt.Printf("Current version: %.2f, Age: %d years\n", version, age)
	fmt.Printf("Language name: %s, Is compiled: %t\n", language, isCompiled)
	fmt.Printf("Project: %s, Difficulty: %s, Open Source: %t\n", projectName, difficulty, isOpenSource)
	fmt.Printf("Multiple values: x=%d, y=%d, z=%d\n", x, y, z)
}

// DataTypesExamples demonstrates Go's built-in data types
func DataTypesExamples() {
	fmt.Println("\n=== Data Types Examples ===")

	// Numeric types - integers
	var smallInt int8 = 127                // 8-bit signed integer (-128 to 127)
	var regularInt int = 42                // Platform dependent (32 or 64 bit)
	var bigInt int64 = 9223372036854775807 // 64-bit signed integer
	var unsignedInt uint = 42              // Unsigned integer (0 to max positive)

	// Numeric types - floating point
	var pi float32 = 3.14159                  // 32-bit floating point
	var precisePI float64 = 3.141592653589793 // 64-bit floating point (default)

	// String type
	var greeting string = "Hello, Backend Engineers!"
	var multiLine = `This is a
multi-line string
using backticks`

	// Boolean type
	var isLearning bool = true
	var isExpert bool = false

	// Character type (rune is an alias for int32)
	var firstLetter rune = 'G' // Unicode code point
	var emoji rune = '🚀'       // Supports Unicode emojis

	// Byte type (alias for uint8)
	var byteValue byte = 65 // ASCII value for 'A'

	// Complex numbers (useful for mathematical calculations)
	var complexNum complex64 = 3 + 4i
	var preciseComplex complex128 = 3.14 + 2.71i

	// Print all data types with their values and types
	fmt.Printf("Integers: int8=%d, int=%d, int64=%d, uint=%d\n", smallInt, regularInt, bigInt, unsignedInt)
	fmt.Printf("Floats: float32=%.5f, float64=%.15f\n", pi, precisePI)
	fmt.Printf("String: %s\n", greeting)
	fmt.Printf("Multi-line string:\n%s\n", multiLine)
	fmt.Printf("Booleans: isLearning=%t, isExpert=%t\n", isLearning, isExpert)
	fmt.Printf("Characters: letter=%c (%d), emoji=%c (%d)\n", firstLetter, firstLetter, emoji, emoji)
	fmt.Printf("Byte: %d (char: %c)\n", byteValue, byteValue)
	fmt.Printf("Complex: %v, %v\n", complexNum, preciseComplex)
}

// ConstantsExamples demonstrates constant declarations in Go
func ConstantsExamples() {
	fmt.Println("\n=== Constants Examples ===")

	// Individual constant declarations
	const ServerPort = 8080                     // Untyped constant
	const DatabaseURL string = "localhost:5432" // Typed constant
	const MaxRetries = 3                        // Numeric constant

	// Grouped constant declarations
	const (
		// HTTP status codes commonly used in backend development
		StatusOK                  = 200
		StatusBadRequest          = 400
		StatusUnauthorized        = 401
		StatusNotFound            = 404
		StatusInternalServerError = 500

		// Configuration constants
		AppName    = "GoLearning Backend"
		AppVersion = "1.0.0"
		Debug      = true
	)

	// iota: Go's constant generator for creating enumerated constants
	const (
		// User roles in a backend system
		RoleGuest      int = iota // 0
		RoleUser                  // 1
		RoleAdmin                 // 2
		RoleSuperAdmin            // 3
	)

	// iota with expressions
	const (
		_  = iota             // Skip first value
		KB = 1 << (10 * iota) // 1024 (2^10)
		MB                    // 1048576 (2^20)
		GB                    // 1073741824 (2^30)
	)

	fmt.Printf("Server Configuration: Port=%d, Database=%s, Max Retries=%d\n",
		ServerPort, DatabaseURL, MaxRetries)
	fmt.Printf("HTTP Status Codes: OK=%d, Not Found=%d, Server Error=%d\n",
		StatusOK, StatusNotFound, StatusInternalServerError)
	fmt.Printf("App Info: %s v%s (Debug: %t)\n", AppName, AppVersion, Debug)
	fmt.Printf("User Roles: Guest=%d, User=%d, Admin=%d, SuperAdmin=%d\n",
		RoleGuest, RoleUser, RoleAdmin, RoleSuperAdmin)
	fmt.Printf("Storage Units: KB=%d, MB=%d, GB=%d\n", KB, MB, GB)
}
