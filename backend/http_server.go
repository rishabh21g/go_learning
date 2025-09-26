// Package backend demonstrates backend engineering concepts in Go
// This package covers HTTP servers, middleware, routing, and API development
package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HTTPServerExamples demonstrates basic HTTP server creation and handling
func HTTPServerExamples() {
	fmt.Println("=== HTTP Server Examples ===")
	fmt.Println("Note: This is a demonstration of HTTP server code.")
	fmt.Println("In a real application, you would run the server in a separate goroutine.")

	// Create a new HTTP server multiplexer (router)
	mux := http.NewServeMux()

	// Register handlers for different routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/health", healthHandler)
	mux.HandleFunc("/api/users", usersHandler)
	mux.HandleFunc("/api/users/", userByIDHandler) // Note the trailing slash for path parameters

	// Create server with configuration
	server := &http.Server{
		Addr:           ":8080",
		Handler:        loggingMiddleware(mux), // Wrap with middleware
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	fmt.Printf("Server configuration:\n")
	fmt.Printf("  Address: %s\n", server.Addr)
	fmt.Printf("  Read Timeout: %v\n", server.ReadTimeout)
	fmt.Printf("  Write Timeout: %v\n", server.WriteTimeout)
	fmt.Printf("  Max Header Bytes: %d\n", server.MaxHeaderBytes)

	// In a real application, you would start the server like this:
	// log.Fatal(server.ListenAndServe())

	fmt.Println("\nAPI Endpoints available:")
	fmt.Println("  GET  /              - Home page")
	fmt.Println("  GET  /api/health    - Health check")
	fmt.Println("  GET  /api/users     - List all users")
	fmt.Println("  POST /api/users     - Create new user")
	fmt.Println("  GET  /api/users/{id} - Get user by ID")
}

// homeHandler handles requests to the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header
	w.Header().Set("Content-Type", "text/html")

	// Write response
	html := `
	<html>
		<head><title>Go Learning Backend</title></head>
		<body>
			<h1>Welcome to Go Backend Development!</h1>
			<p>This is a sample HTTP server built with Go.</p>
			<ul>
				<li><a href="/api/health">Health Check</a></li>
				<li><a href="/api/users">Users API</a></li>
			</ul>
		</body>
	</html>`

	w.Write([]byte(html))
}

// healthHandler provides a health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create health response
	health := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
		Uptime:    "24h30m",
		Services: map[string]string{
			"database": "connected",
			"cache":    "operational",
			"queue":    "running",
		},
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// usersHandler handles requests to /api/users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r)
	case http.MethodPost:
		handleCreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetUsers handles GET /api/users
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	// Simulate database query
	users := []User{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", Role: "admin"},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com", Role: "user"},
		{ID: 3, Name: "Charlie Brown", Email: "charlie@example.com", Role: "user"},
	}

	// Create response
	response := UsersResponse{
		Users: users,
		Total: len(users),
		Page:  1,
		Limit: 10,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleCreateUser handles POST /api/users
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var createReq CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate request
	if createReq.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if createReq.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Create new user (simulate database insert)
	newUser := User{
		ID:    generateUserID(),
		Name:  createReq.Name,
		Email: createReq.Email,
		Role:  "user", // Default role
	}

	// Return created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// userByIDHandler handles requests to /api/users/{id}
func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	if path == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleGetUserByID(w, r, userID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetUserByID handles GET /api/users/{id}
func handleGetUserByID(w http.ResponseWriter, r *http.Request, userID int) {
	// Simulate database lookup
	users := map[int]User{
		1: {ID: 1, Name: "Alice Johnson", Email: "alice@example.com", Role: "admin"},
		2: {ID: 2, Name: "Bob Smith", Email: "bob@example.com", Role: "user"},
		3: {ID: 3, Name: "Charlie Brown", Email: "charlie@example.com", Role: "user"},
	}

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// generateUserID generates a simple user ID (in real app, use database auto-increment)
func generateUserID() int {
	return int(time.Now().Unix() % 10000)
}

// Middleware functions

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request
		log.Printf("Started %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Call next handler
		next.ServeHTTP(w, r)

		// Log completion
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	})
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// authMiddleware provides basic authentication
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check for Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token (simplified - in real app, verify JWT or check database)
		if token != "valid-token-123" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Response types

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Version   string            `json:"version"`
	Uptime    string            `json:"uptime"`
	Services  map[string]string `json:"services"`
}

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UsersResponse represents the response for listing users
type UsersResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

// MiddlewareExamples demonstrates how to compose middleware
func MiddlewareExamples() {
	fmt.Println("\n=== Middleware Examples ===")

	// Create a basic handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "Hello from protected endpoint!",
			"method":  r.Method,
			"path":    r.URL.Path,
		}
		json.NewEncoder(w).Encode(response)
	})

	// Compose middleware (order matters!)
	protectedHandler := loggingMiddleware(
		corsMiddleware(
			authMiddleware(handler),
		),
	)

	fmt.Println("Middleware chain created:")
	fmt.Println("  1. Logging Middleware - logs all requests")
	fmt.Println("  2. CORS Middleware - adds cross-origin headers")
	fmt.Println("  3. Auth Middleware - validates Bearer tokens")
	fmt.Println("  4. Final Handler - processes the request")

	// Example of how middleware wraps the handler
	fmt.Println("\nRequest flow:")
	fmt.Println("  Request → Logging → CORS → Auth → Handler → Auth → CORS → Logging → Response")

	// In a real server, you would register this handler:
	// http.Handle("/api/protected", protectedHandler)

	// Demonstrate different middleware combinations
	fmt.Println("\nDifferent middleware combinations:")

	// Public endpoint (only logging and CORS)
	publicHandler := loggingMiddleware(corsMiddleware(handler))
	fmt.Println("  Public endpoint: Logging + CORS")

	// Admin endpoint (all middleware + rate limiting)
	adminHandler := loggingMiddleware(
		corsMiddleware(
			authMiddleware(
				rateLimitMiddleware(handler),
			),
		),
	)
	fmt.Println("  Admin endpoint: Logging + CORS + Auth + Rate Limiting")

	// Store handlers to avoid "declared but not used" error
	_ = protectedHandler
	_ = publicHandler
	_ = adminHandler
}

// rateLimitMiddleware provides basic rate limiting
func rateLimitMiddleware(next http.Handler) http.Handler {
	// Simple in-memory rate limiter (in production, use Redis or similar)
	requests := make(map[string][]time.Time)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		now := time.Now()

		// Clean old requests (older than 1 minute)
		if reqTimes, exists := requests[clientIP]; exists {
			var validRequests []time.Time
			for _, reqTime := range reqTimes {
				if now.Sub(reqTime) <= time.Minute {
					validRequests = append(validRequests, reqTime)
				}
			}
			requests[clientIP] = validRequests
		}

		// Check rate limit (max 10 requests per minute)
		if len(requests[clientIP]) >= 10 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Add current request
		requests[clientIP] = append(requests[clientIP], now)

		next.ServeHTTP(w, r)
	})
}
