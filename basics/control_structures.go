// Package basics - control structures and flow control in Go
package basics

import (
	"fmt"
	"strings"
)

// ConditionalExamples demonstrates if/else statements and switch cases
func ConditionalExamples() {
	fmt.Println("\n=== Conditional Statements Examples ===")

	// Basic if statement
	serverStatus := "running"
	if serverStatus == "running" {
		fmt.Println("✅ Server is operational")
	}

	// If with initialization statement (common pattern in Go)
	if port := 8080; port > 1024 {
		fmt.Printf("✅ Server running on port %d (above privileged range)\n", port)
	}

	// If-else chain for HTTP status handling
	statusCode := 404
	if statusCode >= 200 && statusCode < 300 {
		fmt.Println("✅ Success response")
	} else if statusCode >= 400 && statusCode < 500 {
		fmt.Println("❌ Client error")
	} else if statusCode >= 500 {
		fmt.Println("💥 Server error")
	} else {
		fmt.Println("ℹ️  Informational response")
	}

	// Switch statement - more efficient than if-else chains
	httpMethod := "POST"
	switch httpMethod {
	case "GET":
		fmt.Println("📖 Reading data")
	case "POST":
		fmt.Println("📝 Creating new resource")
	case "PUT":
		fmt.Println("✏️  Updating existing resource")
	case "DELETE":
		fmt.Println("🗑️  Deleting resource")
	default:
		fmt.Println("❓ Unknown HTTP method")
	}

	// Switch with multiple cases
	contentType := "application/json"
	switch contentType {
	case "application/json", "application/xml":
		fmt.Println("📊 Structured data format")
	case "text/plain", "text/html":
		fmt.Println("📄 Text-based format")
	case "image/jpeg", "image/png", "image/gif":
		fmt.Println("🖼️  Image format")
	default:
		fmt.Println("📎 Other format")
	}

	// Switch without expression (replaces if-else)
	errorCount := 15
	switch {
	case errorCount == 0:
		fmt.Println("🎉 No errors!")
	case errorCount < 10:
		fmt.Println("⚠️  Few errors, monitoring required")
	case errorCount < 50:
		fmt.Println("🚨 Moderate errors, investigation needed")
	default:
		fmt.Println("💥 Critical error count, immediate action required")
	}
}

// LoopExamples demonstrates different types of loops in Go
func LoopExamples() {
	fmt.Println("\n=== Loop Examples ===")

	// Basic for loop (C-style)
	fmt.Println("📊 Processing requests:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Request %d processed\n", i)
	}

	// While-style loop (condition only)
	fmt.Println("\n⏳ Server startup sequence:")
	attempts := 0
	maxAttempts := 3
	for attempts < maxAttempts {
		attempts++
		fmt.Printf("  Startup attempt %d/%d\n", attempts, maxAttempts)
	}

	// Infinite loop with break (common in server applications)
	fmt.Println("\n🔄 Server listening loop simulation:")
	requestCount := 0
	for {
		requestCount++
		fmt.Printf("  Handling request #%d\n", requestCount)

		// Break condition (simulating server shutdown)
		if requestCount >= 3 {
			fmt.Println("  🛑 Server shutdown initiated")
			break
		}
	}

	// For-range with slices (very common in Go)
	fmt.Println("\n📝 Processing user list:")
	users := []string{"alice", "bob", "charlie", "diana"}
	for index, user := range users {
		fmt.Printf("  User %d: %s\n", index+1, strings.Title(user))
	}

	// For-range with maps (key-value iteration)
	fmt.Println("\n🔧 Server configuration:")
	config := map[string]interface{}{
		"port":     8080,
		"debug":    true,
		"max_conn": 100,
		"timeout":  "30s",
	}
	for key, value := range config {
		fmt.Printf("  %s: %v\n", key, value)
	}

	// For-range with strings (iterates over runes)
	fmt.Println("\n🔤 Processing text characters:")
	text := "Go🚀"
	for i, char := range text {
		fmt.Printf("  Position %d: %c (Unicode: %d)\n", i, char, char)
	}

	// Using continue to skip iterations
	fmt.Println("\n🔍 Filtering valid user IDs:")
	userIDs := []int{1, -1, 5, 0, 8, -3, 12}
	for _, id := range userIDs {
		if id <= 0 {
			fmt.Printf("  Skipping invalid ID: %d\n", id)
			continue
		}
		fmt.Printf("  ✅ Valid user ID: %d\n", id)
	}
}

// CollectionsExamples demonstrates arrays, slices, and maps
func CollectionsExamples() {
	fmt.Println("\n=== Collections Examples ===")

	// Arrays (fixed size, rarely used directly)
	fmt.Println("📦 Arrays (fixed size):")
	var httpMethods [4]string
	httpMethods[0] = "GET"
	httpMethods[1] = "POST"
	httpMethods[2] = "PUT"
	httpMethods[3] = "DELETE"
	fmt.Printf("  HTTP Methods: %v (Length: %d)\n", httpMethods, len(httpMethods))

	// Array literal initialization
	statusCodes := [5]int{200, 201, 400, 404, 500}
	fmt.Printf("  Status Codes: %v\n", statusCodes)

	// Slices (dynamic arrays, most commonly used)
	fmt.Println("\n📋 Slices (dynamic arrays):")

	// Creating slices
	var endpoints []string                                    // nil slice
	endpoints = append(endpoints, "/api/users", "/api/posts") // adding elements
	endpoints = append(endpoints, "/api/comments")
	fmt.Printf("  API Endpoints: %v (Length: %d, Capacity: %d)\n",
		endpoints, len(endpoints), cap(endpoints))

	// Slice literal
	middleware := []string{"cors", "auth", "logging", "ratelimit"}
	fmt.Printf("  Middleware: %v\n", middleware)

	// Slice operations
	fmt.Println("  Slice operations:")
	fmt.Printf("    First middleware: %s\n", middleware[0])
	fmt.Printf("    Last middleware: %s\n", middleware[len(middleware)-1])
	fmt.Printf("    Middle two: %v\n", middleware[1:3])   // slicing
	fmt.Printf("    From index 2: %v\n", middleware[2:])  // slice to end
	fmt.Printf("    Up to index 2: %v\n", middleware[:2]) // slice from start

	// Making slices with specific capacity
	requestQueue := make([]string, 0, 10) // length 0, capacity 10
	requestQueue = append(requestQueue, "req1", "req2", "req3")
	fmt.Printf("  Request Queue: %v (Len: %d, Cap: %d)\n",
		requestQueue, len(requestQueue), cap(requestQueue))

	// Maps (key-value pairs, like dictionaries/hashmaps)
	fmt.Println("\n🗂️  Maps (key-value pairs):")

	// Creating and initializing maps
	userRoles := make(map[string]string)
	userRoles["admin"] = "full_access"
	userRoles["user"] = "read_write"
	userRoles["guest"] = "read_only"
	fmt.Printf("  User Roles: %v\n", userRoles)

	// Map literal initialization
	serverStats := map[string]int{
		"active_connections": 150,
		"requests_per_sec":   45,
		"error_count":        2,
		"uptime_hours":       72,
	}
	fmt.Printf("  Server Stats: %v\n", serverStats)

	// Map operations
	fmt.Println("  Map operations:")

	// Checking if key exists
	if role, exists := userRoles["admin"]; exists {
		fmt.Printf("    Admin role: %s ✅\n", role)
	}

	if _, exists := userRoles["superuser"]; !exists {
		fmt.Println("    Superuser role: not found ❌")
	}

	// Iterating over maps
	fmt.Println("    All server statistics:")
	for metric, value := range serverStats {
		fmt.Printf("      %s: %d\n", metric, value)
	}

	// Deleting from maps
	delete(userRoles, "guest")
	fmt.Printf("  After removing guest: %v\n", userRoles)

	// Nested collections (common in backend development)
	fmt.Println("\n🏗️  Nested Collections:")

	// Slice of maps (like JSON array of objects)
	users := []map[string]interface{}{
		{"id": 1, "name": "Alice", "role": "admin", "active": true},
		{"id": 2, "name": "Bob", "role": "user", "active": true},
		{"id": 3, "name": "Charlie", "role": "user", "active": false},
	}

	fmt.Println("  User database:")
	for i, user := range users {
		fmt.Printf("    User %d: %v\n", i+1, user)
	}

	// Map of slices (like grouped data)
	usersByRole := map[string][]string{
		"admin": {"alice", "david"},
		"user":  {"bob", "charlie", "eve"},
		"guest": {"anonymous1", "anonymous2"},
	}

	fmt.Println("  Users grouped by role:")
	for role, userList := range usersByRole {
		fmt.Printf("    %s: %v\n", role, userList)
	}
}
