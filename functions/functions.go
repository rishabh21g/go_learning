// Package functions demonstrates Go function syntax and patterns
// This package covers function declarations, parameters, return values, and advanced patterns
package functions

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// BasicFunctionExamples demonstrates simple function declarations and calls
func BasicFunctionExamples() {
	fmt.Println("=== Basic Function Examples ===")

	// Simple function call
	greeting := greetUser("Backend Developer")
	fmt.Println(greeting)

	// Function with multiple parameters
	result := addNumbers(25, 17)
	fmt.Printf("25 + 17 = %d\n", result)

	// Function with multiple return values
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)

	// Function returning error (Go's error handling pattern)
	safeResult, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Safe division result: %.2f\n", safeResult)
	}
}

// greetUser is a simple function that takes a string parameter and returns a string
func greetUser(role string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go backend development.", role)
}

// addNumbers demonstrates a function with multiple parameters of the same type
func addNumbers(a, b int) int {
	return a + b
}

// divide demonstrates multiple return values (common in Go)
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// safeDivide demonstrates Go's error handling pattern
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// AdvancedFunctionExamples demonstrates advanced function features
func AdvancedFunctionExamples() {
	fmt.Println("\n=== Advanced Function Examples ===")

	// Named return values
	area, perimeter := calculateRectangle(5, 3)
	fmt.Printf("Rectangle (5x3): Area = %d, Perimeter = %d\n", area, perimeter)

	// Variadic functions (accepting variable number of arguments)
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)

	// Variadic with slice
	numbers := []int{10, 20, 30}
	sliceTotal := sum(numbers...) // spread operator
	fmt.Printf("Sum of slice [10,20,30] = %d\n", sliceTotal)

	// Function as variable
	var operation func(int, int) int = multiply
	result := operation(4, 7)
	fmt.Printf("4 × 7 = %d (using function variable)\n", result)

	// Anonymous function (lambda)
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 6 = %d\n", square(6))

	// Immediately invoked function expression (IIFE)
	message := func(name string) string {
		return fmt.Sprintf("Processing user: %s", name)
	}("Alice")
	fmt.Println(message)
}

// calculateRectangle demonstrates named return values
func calculateRectangle(length, width int) (area int, perimeter int) {
	area = length * width // named returns are automatically returned
	perimeter = 2 * (length + width)
	return // naked return - returns named values
}

// sum is a variadic function that accepts any number of integers
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// multiply is used to demonstrate function variables
func multiply(a, b int) int {
	return a * b
}

// HigherOrderFunctions demonstrates functions that work with other functions
func HigherOrderFunctions() {
	fmt.Println("\n=== Higher-Order Functions ===")

	// Function that takes another function as parameter
	numbers := []int{1, 2, 3, 4, 5}

	// Using different operations
	doubled := applyOperation(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Doubled: %v\n", doubled)

	squared := applyOperation(numbers, func(x int) int { return x * x })
	fmt.Printf("Squared: %v\n", squared)

	// Function returning another function (closure)
	multiplier := createMultiplier(3)
	fmt.Printf("3 × 7 = %d (using closure)\n", multiplier(7))

	// Practical example: middleware pattern (common in web servers)
	handler := createLoggingMiddleware(businessLogic)
	handler("Process user data")
}

// applyOperation applies a function to each element in a slice
func applyOperation(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation(num)
	}
	return result
}

// createMultiplier returns a function that multiplies by a given factor (closure)
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// businessLogic represents some business logic in our application
func businessLogic(data string) {
	fmt.Printf("  Executing business logic with: %s\n", data)
}

// createLoggingMiddleware creates a wrapper function that adds logging
func createLoggingMiddleware(next func(string)) func(string) {
	return func(data string) {
		fmt.Printf("  [LOG] Starting operation at %s\n", time.Now().Format("15:04:05"))
		next(data)
		fmt.Printf("  [LOG] Operation completed\n")
	}
}

// ErrorHandlingPatterns demonstrates common Go error handling patterns
func ErrorHandlingPatterns() {
	fmt.Println("\n=== Error Handling Patterns ===")

	// Pattern 1: Simple error check
	if err := validateEmail("invalid-email"); err != nil {
		fmt.Printf("Email validation failed: %v\n", err)
	}

	if err := validateEmail("user@example.com"); err != nil {
		fmt.Printf("Email validation failed: %v\n", err)
	} else {
		fmt.Println("✅ Email is valid")
	}

	// Pattern 2: Early return pattern
	user, err := createUser("john_doe", "john@example.com")
	if err != nil {
		fmt.Printf("User creation failed: %v\n", err)
		return
	}
	fmt.Printf("✅ User created: %v\n", user)

	// Pattern 3: Error wrapping (Go 1.13+)
	if err := processUserData(user); err != nil {
		fmt.Printf("Processing failed: %v\n", err)
	}

	// Pattern 4: Multiple error handling
	results, warnings, err := batchProcess([]string{"item1", "item2", "item3"})
	if err != nil {
		fmt.Printf("Batch processing failed: %v\n", err)
		return
	}
	if len(warnings) > 0 {
		fmt.Printf("⚠️  Warnings: %v\n", warnings)
	}
	fmt.Printf("✅ Processing results: %v\n", results)
}

// User represents a user in our system
type User struct {
	Username string
	Email    string
	ID       int
}

// validateEmail validates an email address format
func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("email must contain @ symbol")
	}
	if !strings.Contains(email, ".") {
		return errors.New("email must contain domain extension")
	}
	return nil
}

// createUser creates a new user with validation
func createUser(username, email string) (*User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if err := validateEmail(email); err != nil {
		return nil, fmt.Errorf("invalid email: %w", err) // error wrapping
	}

	return &User{
		Username: username,
		Email:    email,
		ID:       generateUserID(),
	}, nil
}

// generateUserID generates a unique user ID (simplified)
func generateUserID() int {
	return int(time.Now().Unix() % 10000)
}

// processUserData processes user data and may return wrapped errors
func processUserData(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	// Simulate some processing that might fail
	if len(user.Username) < 3 {
		return fmt.Errorf("processing failed for user %s: %w",
			user.Username, errors.New("username too short"))
	}

	fmt.Printf("  ✅ Successfully processed user: %s\n", user.Username)
	return nil
}

// batchProcess demonstrates handling multiple types of results
func batchProcess(items []string) ([]string, []string, error) {
	if len(items) == 0 {
		return nil, nil, errors.New("no items to process")
	}

	var results []string
	var warnings []string

	for i, item := range items {
		if item == "" {
			warnings = append(warnings, fmt.Sprintf("empty item at index %d", i))
			continue
		}

		// Simulate processing
		processed := fmt.Sprintf("processed_%s", item)
		results = append(results, processed)

		// Simulate conditional warning
		if len(item) < 4 {
			warnings = append(warnings, fmt.Sprintf("short item name: %s", item))
		}
	}

	return results, warnings, nil
}

// MethodExamples demonstrates methods (functions with receivers)
func MethodExamples() {
	fmt.Println("\n=== Method Examples ===")

	// Creating instances
	server := &Server{
		Name: "WebServer-1",
		Port: 8080,
	}

	// Calling methods
	server.Start()
	server.HandleRequest("GET /api/users")
	server.SetPort(9090)
	status := server.GetStatus()
	fmt.Printf("Server status: %s\n", status)
	server.Stop()
}

// Server represents a web server
type Server struct {
	Name    string
	Port    int
	Running bool
}

// Start starts the server (method with pointer receiver)
func (s *Server) Start() {
	s.Running = true
	fmt.Printf("  🚀 Server %s started on port %d\n", s.Name, s.Port)
}

// Stop stops the server
func (s *Server) Stop() {
	s.Running = false
	fmt.Printf("  🛑 Server %s stopped\n", s.Name)
}

// HandleRequest handles an incoming request
func (s *Server) HandleRequest(request string) {
	if !s.Running {
		fmt.Printf("  ❌ Cannot handle request: server is not running\n")
		return
	}
	fmt.Printf("  📨 %s handling: %s\n", s.Name, request)
}

// SetPort updates the server port (requires pointer receiver to modify)
func (s *Server) SetPort(port int) {
	s.Port = port
	fmt.Printf("  🔧 %s port updated to %d\n", s.Name, port)
}

// GetStatus returns server status (value receiver is fine for read-only)
func (s Server) GetStatus() string {
	if s.Running {
		return fmt.Sprintf("Running on port %d", s.Port)
	}
	return "Stopped"
}
