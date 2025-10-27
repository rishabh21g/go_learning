// Package main demonstrates a comprehensive Go learning application
// This is the entry point that showcases all Go concepts covered in this learning project
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rishabh21g/go_learning/backend"
	"github.com/rishabh21g/go_learning/basics"
	"github.com/rishabh21g/go_learning/concurrency"
	"github.com/rishabh21g/go_learning/functions"
	"github.com/rishabh21g/go_learning/structs"
)

// main is the entry point of the Go learning application
func main() {
	// Display welcome message
	displayWelcome()

	// Interactive menu system
	for {
		displayMenu()
		choice := getUserInput("Enter your choice (1-7, or 0 to exit): ")

		if choice == "0" {
			fmt.Println("\n👋 Thank you for learning Go! Happy coding!")
			break
		}

		executeChoice(choice)

		// Wait for user to continue
		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

// displayWelcome shows the application header
func displayWelcome() {
	fmt.Println("🚀 ============================================== 🚀")
	fmt.Println("   Welcome to Go Learning for Backend Engineers")
	fmt.Println("🚀 ============================================== 🚀")
	fmt.Println()
	fmt.Println("This interactive application demonstrates:")
	fmt.Println("✅ Go syntax and language fundamentals")
	fmt.Println("✅ Functions and error handling patterns")
	fmt.Println("✅ Structs, interfaces, and composition")
	fmt.Println("✅ Backend engineering concepts")
	fmt.Println("✅ Concurrency with goroutines and channels")
	fmt.Println("✅ Real-world patterns and best practices")
	fmt.Println()
}

// displayMenu shows the main menu options
func displayMenu() {
	fmt.Println("📚 =========================")
	fmt.Println("   LEARNING MENU")
	fmt.Println("📚 =========================")
	fmt.Println("1. 🏗️  Basic Syntax & Data Types")
	fmt.Println("2. 🔄 Control Structures & Collections")
	fmt.Println("3. ⚙️  Functions & Error Handling")
	fmt.Println("4. 🏛️  Structs & Interfaces")
	fmt.Println("5. 🌐 Backend HTTP Server")
	fmt.Println("6. 🚦 Concurrency & Goroutines")
	fmt.Println("7. 📖 All Examples (Full Demo)")
	fmt.Println("0. 🚪 Exit")
	fmt.Println("📚 =========================")
}

// getUserInput prompts the user and returns their input
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// executeChoice runs the selected learning module
func executeChoice(choice string) {
	fmt.Println() // Add spacing

	switch choice {
	case "1":
		runBasicSyntax()
	case "2":
		runControlStructures()
	case "3":
		runFunctions()
	case "4":
		runStructsInterfaces()
	case "5":
		runBackendConcepts()
	case "6":
		runConcurrency()
	case "7":
		runAllExamples()
	default:
		fmt.Println("❌ Invalid choice. Please select a number from 0-7.")
	}
}

// runBasicSyntax demonstrates basic Go syntax
func runBasicSyntax() {
	printSectionHeader("BASIC SYNTAX & DATA TYPES")

	// Variable examples
	basics.VariableExamples()

	// Data types examples
	basics.DataTypesExamples()

	// Constants examples
	basics.ConstantsExamples()

	printSectionFooter("Completed: Basic Syntax & Data Types")
}

// runControlStructures demonstrates control structures and collections
func runControlStructures() {
	printSectionHeader("CONTROL STRUCTURES & COLLECTIONS")

	// Conditional statements
	basics.ConditionalExamples()

	// Loops
	basics.LoopExamples()

	// Collections (arrays, slices, maps)
	basics.CollectionsExamples()

	printSectionFooter("Completed: Control Structures & Collections")
}

// runFunctions demonstrates functions and error handling
func runFunctions() {
	printSectionHeader("FUNCTIONS & ERROR HANDLING")

	// Basic functions
	functions.BasicFunctionExamples()

	// Advanced functions
	functions.AdvancedFunctionExamples()

	// Higher-order functions
	functions.HigherOrderFunctions()

	// Error handling patterns
	functions.ErrorHandlingPatterns()

	// Methods
	functions.MethodExamples()

	printSectionFooter("Completed: Functions & Error Handling")
}

// runStructsInterfaces demonstrates structs and interfaces
func runStructsInterfaces() {
	printSectionHeader("STRUCTS & INTERFACES")

	// Struct examples
	structs.StructExamples()

	// Interface examples
	structs.InterfaceExamples()

	// Advanced patterns
	structs.AdvancedPatterns()

	// Composition examples
	structs.CompositionExamples()

	printSectionFooter("Completed: Structs & Interfaces")
}

// runBackendConcepts demonstrates backend engineering concepts
func runBackendConcepts() {
	printSectionHeader("BACKEND ENGINEERING CONCEPTS")

	// HTTP server examples
	backend.HTTPServerExamples()

	// Middleware examples
	backend.MiddlewareExamples()

	// Offer to start a real server
	fmt.Println("\n🤔 Would you like to see a practical demonstration?")
	startServer := getUserInput("Start a real HTTP server? (y/n): ")

	if strings.ToLower(startServer) == "y" || strings.ToLower(startServer) == "yes" {
		fmt.Println("\n🚀 Starting demonstration HTTP server...")
		fmt.Println("Note: In this learning environment, we'll simulate the server.")
		fmt.Println("In a real application, you would access:")
		fmt.Println("  • http://localhost:8080/ - Home page")
		fmt.Println("  • http://localhost:8080/api/health - Health check")
		fmt.Println("  • http://localhost:8080/api/users - Users API")
		fmt.Println("\n⚠️  The server would run with: go run main.go")
		fmt.Println("   Then use curl or a browser to test the endpoints.")
	}

	printSectionFooter("Completed: Backend Engineering Concepts")
}

// runConcurrency demonstrates concurrency concepts
func runConcurrency() {
	printSectionHeader("CONCURRENCY & GOROUTINES")

	// Ask user which examples to run (some are time-consuming)
	fmt.Println("Concurrency examples available:")
	fmt.Println("1. Basic Goroutines")
	fmt.Println("2. WaitGroups & Synchronization")
	fmt.Println("3. Channels & Communication")
	fmt.Println("4. Select Statements")
	fmt.Println("5. Producer-Consumer Pattern")
	fmt.Println("6. Context & Cancellation")
	fmt.Println("7. Pipeline Pattern")
	fmt.Println("8. Worker Pool Pattern")
	fmt.Println("9. All Concurrency Examples")

	choice := getUserInput("Which examples would you like to see? (1-9): ")

	switch choice {
	case "1":
		concurrency.GoroutineExamples()
	case "2":
		concurrency.WaitGroupExamples()
	case "3":
		concurrency.ChannelExamples()
	case "4":
		concurrency.SelectExamples()
	case "5":
		concurrency.ProducerConsumerPattern()
	case "6":
		concurrency.ContextExamples()
	case "7":
		concurrency.PipelinePattern()
	case "8":
		concurrency.WorkerPoolPattern()
	case "9":
		runAllConcurrencyExamples()
	default:
		fmt.Println("❌ Invalid choice. Running basic goroutine examples...")
		concurrency.GoroutineExamples()
	}

	printSectionFooter("Completed: Concurrency & Goroutines")
}

// runAllConcurrencyExamples runs all concurrency examples
func runAllConcurrencyExamples() {
	fmt.Println("🔄 Running all concurrency examples...")
	fmt.Println("⚠️  This may take a few minutes due to timing demonstrations.")

	concurrency.GoroutineExamples()
	concurrency.WaitGroupExamples()
	concurrency.ChannelExamples()
	concurrency.SelectExamples()
	concurrency.ProducerConsumerPattern()
	concurrency.ContextExamples()
	concurrency.PipelinePattern()
	concurrency.WorkerPoolPattern()
}

// runAllExamples runs all examples in sequence
func runAllExamples() {
	printSectionHeader("COMPLETE GO LEARNING DEMONSTRATION")

	fmt.Println("🎯 This will run ALL examples from all modules.")
	fmt.Println("⏰ Estimated time: 5-10 minutes")
	fmt.Println("🔄 Some examples include timing demonstrations.")

	proceed := getUserInput("Continue with full demonstration? (y/n): ")
	if strings.ToLower(proceed) != "y" && strings.ToLower(proceed) != "yes" {
		fmt.Println("Demonstration cancelled.")
		return
	}

	// Progress tracking
	modules := []struct {
		name string
		fn   func()
	}{
		{"Basic Syntax", runBasicSyntaxDemo},
		{"Control Structures", runControlStructuresDemo},
		{"Functions", runFunctionsDemo},
		{"Structs & Interfaces", runStructsInterfacesDemo},
		{"Backend Concepts", runBackendConceptsDemo},
		{"Concurrency", runConcurrencyDemo},
	}

	totalModules := len(modules)

	for i, module := range modules {
		fmt.Printf("\n🔄 Progress: %d/%d - %s\n", i+1, totalModules, module.name)
		module.fn()

		// Progress bar
		progress := float64(i+1) / float64(totalModules) * 100
		fmt.Printf("📊 Overall Progress: %.0f%% complete\n", progress)

		if i < totalModules-1 {
			fmt.Println("⏳ Moving to next module in 2 seconds...")
			// time.Sleep(2 * time.Second) // Commented out for faster demo
		}
	}

	printSectionFooter("🎉 COMPLETE DEMONSTRATION FINISHED!")
	fmt.Println("🎯 Congratulations! You've seen all major Go concepts.")
	fmt.Println("📚 Next steps:")
	fmt.Println("   • Build your own Go projects")
	fmt.Println("   • Explore the Go standard library")
	fmt.Println("   • Learn popular Go frameworks (Gin, Echo, etc.)")
	fmt.Println("   • Practice with Go modules and dependency management")
	fmt.Println("   • Study Go's testing framework")
}

// Demo functions (simplified versions for full demonstration)
func runBasicSyntaxDemo() {
	basics.VariableExamples()
	basics.DataTypesExamples()
}

func runControlStructuresDemo() {
	basics.ConditionalExamples()
	basics.LoopExamples()
}

func runFunctionsDemo() {
	functions.BasicFunctionExamples()
	functions.ErrorHandlingPatterns()
}

func runStructsInterfacesDemo() {
	structs.StructExamples()
	structs.InterfaceExamples()
}

func runBackendConceptsDemo() {
	backend.HTTPServerExamples()
}

func runConcurrencyDemo() {
	concurrency.GoroutineExamples()
	concurrency.ChannelExamples()
}

// Utility functions for formatting

// printSectionHeader prints a formatted section header
func printSectionHeader(title string) {
	border := strings.Repeat("=", len(title)+6)
	fmt.Printf("\n🎯 %s\n", border)
	fmt.Printf("   %s\n", title)
	fmt.Printf("🎯 %s\n\n", border)
}

// printSectionFooter prints a formatted section footer
func printSectionFooter(message string) {
	fmt.Printf("\n✅ %s\n", message)
	fmt.Println(strings.Repeat("-", len(message)+3))
}

// Additional utility functions for demonstration

// DemoInfo represents information about this learning project
type DemoInfo struct {
	ProjectName   string
	Version       string
	Author        string
	Description   string
	Topics        []string
	Prerequisites []string
}

// GetProjectInfo returns information about this learning project
func GetProjectInfo() DemoInfo {
	return DemoInfo{
		ProjectName: "Go Learning for Backend Engineers",
		Version:     "1.0.0",
		Author:      "Learning Project",
		Description: "Comprehensive Go language learning with focus on backend development",
		Topics: []string{
			"Go Syntax and Fundamentals",
			"Functions and Error Handling",
			"Structs and Interfaces",
			"HTTP Servers and Middleware",
			"Concurrency with Goroutines",
			"Channels and Select Statements",
			"Backend Engineering Patterns",
			"Real-world Go Applications",
		},
		Prerequisites: []string{
			"Basic programming knowledge",
			"Understanding of backend concepts",
			"Familiarity with command line",
			"Go development environment",
		},
	}
}

// printProjectInfo displays project information
func printProjectInfo() {
	info := GetProjectInfo()
	fmt.Printf("📋 Project: %s v%s\n", info.ProjectName, info.Version)
	fmt.Printf("👨‍💻 Author: %s\n", info.Author)
	fmt.Printf("📝 Description: %s\n\n", info.Description)

	fmt.Println("📚 Topics Covered:")
	for i, topic := range info.Topics {
		fmt.Printf("   %d. %s\n", i+1, topic)
	}

	fmt.Println("\n🎯 Prerequisites:")
	for _, prereq := range info.Prerequisites {
		fmt.Printf("   • %s\n", prereq)
	}
}

// validateMenuChoice validates user menu input
func validateMenuChoice(input string) (int, error) {
	choice, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("invalid input: please enter a number")
	}

	if choice < 0 || choice > 7 {
		return -1, fmt.Errorf("choice out of range: please select 0-7")
	}

	return choice, nil
}

// init function runs before main (demonstrates initialization)
func init() {
	// This runs before main() - useful for setup/configuration
	fmt.Println("🔧 Initializing Go Learning Application...")

	// In a real application, you might:
	// - Load configuration files
	// - Set up logging
	// - Initialize database connections
	// - Validate environment variables
}
