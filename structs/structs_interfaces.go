// Package structs demonstrates Go structs, interfaces, and object-oriented patterns
// This package covers struct definitions, methods, interfaces, and composition
package structs

import (
	"encoding/json"
	"fmt"
	"time"
)

// StructExamples demonstrates struct definition and usage
func StructExamples() {
	fmt.Println("=== Struct Examples ===")

	// Creating structs using different methods

	// Method 1: Struct literal with field names
	user1 := User{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		Role:     "admin",
		Active:   true,
		Created:  time.Now(),
	}

	// Method 2: Struct literal with positional values (not recommended)
	user2 := User{2, "bob", "bob@example.com", "user", true, time.Now()}

	// Method 3: Zero value struct and field assignment
	var user3 User
	user3.ID = 3
	user3.Username = "charlie"
	user3.Email = "charlie@example.com"
	user3.Role = "user"
	user3.Active = false
	user3.Created = time.Now()

	// Method 4: Using new() - returns pointer to zero value
	user4 := new(User)
	user4.ID = 4
	user4.Username = "diana"
	user4.Email = "diana@example.com"
	user4.Role = "moderator"
	user4.Active = true
	user4.Created = time.Now()

	fmt.Printf("User 1: %+v\n", user1)
	fmt.Printf("User 2: %+v\n", user2)
	fmt.Printf("User 3: %+v\n", user3)
	fmt.Printf("User 4: %+v\n", *user4) // dereference pointer

	// Working with methods
	fmt.Println("\n--- User Methods ---")
	fmt.Printf("User 1 display name: %s\n", user1.GetDisplayName())
	fmt.Printf("User 1 is admin: %t\n", user1.IsAdmin())

	user1.Deactivate()
	fmt.Printf("User 1 after deactivation: Active = %t\n", user1.Active)

	// Struct embedding (composition)
	fmt.Println("\n--- Struct Embedding ---")
	admin := Admin{
		User: User{
			ID:       5,
			Username: "superadmin",
			Email:    "admin@example.com",
			Role:     "admin",
			Active:   true,
			Created:  time.Now(),
		},
		Permissions: []string{"read", "write", "delete", "manage"},
		Department:  "IT",
	}

	// Can access embedded struct fields directly
	fmt.Printf("Admin username: %s\n", admin.Username)
	fmt.Printf("Admin permissions: %v\n", admin.Permissions)
	fmt.Printf("Admin can manage: %t\n", admin.CanManage())
}

// User represents a user in the system
type User struct {
	ID       int       `json:"id"` // JSON tags for serialization
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created"`
}

// GetDisplayName returns a formatted display name for the user
func (u User) GetDisplayName() string {
	return fmt.Sprintf("%s (%s)", u.Username, u.Role)
}

// IsAdmin checks if the user has admin privileges
func (u User) IsAdmin() bool {
	return u.Role == "admin" || u.Role == "superadmin"
}

// Deactivate deactivates the user (pointer receiver to modify struct)
func (u *User) Deactivate() {
	u.Active = false
}

// Activate activates the user
func (u *User) Activate() {
	u.Active = true
}

// Admin represents an admin user with additional properties
type Admin struct {
	User                 // Embedded struct (composition)
	Permissions []string `json:"permissions"`
	Department  string   `json:"department"`
}

// CanManage checks if admin has management permissions
func (a Admin) CanManage() bool {
	for _, perm := range a.Permissions {
		if perm == "manage" {
			return true
		}
	}
	return false
}

// InterfaceExamples demonstrates Go interfaces
func InterfaceExamples() {
	fmt.Println("\n=== Interface Examples ===")

	// Creating different types that implement the same interface
	var storage DataStorage

	// FileStorage implementation
	fileStorage := &FileStorage{Path: "/data/files"}
	storage = fileStorage
	testStorage(storage, "File Storage")

	// DatabaseStorage implementation
	dbStorage := &DatabaseStorage{
		Host:     "localhost",
		Port:     5432,
		Database: "myapp",
	}
	storage = dbStorage
	testStorage(storage, "Database Storage")

	// MemoryStorage implementation
	memStorage := &MemoryStorage{Data: make(map[string]interface{})}
	storage = memStorage
	testStorage(storage, "Memory Storage")

	// Interface polymorphism - slice of interfaces
	fmt.Println("\n--- Interface Polymorphism ---")
	storages := []DataStorage{fileStorage, dbStorage, memStorage}

	for i, s := range storages {
		fmt.Printf("Storage %d:\n", i+1)
		s.Store("key", fmt.Sprintf("value-%d", i+1))
		value := s.Retrieve("key")
		fmt.Printf("  Retrieved: %v\n", value)
	}
}

// testStorage tests a storage implementation
func testStorage(storage DataStorage, name string) {
	fmt.Printf("\n--- Testing %s ---\n", name)
	storage.Connect()
	storage.Store("user:1", map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	})
	data := storage.Retrieve("user:1")
	fmt.Printf("Retrieved data: %v\n", data)
	storage.Disconnect()
}

// DataStorage interface defines storage operations
type DataStorage interface {
	Connect() error
	Disconnect() error
	Store(key string, value interface{}) error
	Retrieve(key string) interface{}
}

// FileStorage implements DataStorage for file-based storage
type FileStorage struct {
	Path      string
	connected bool
}

func (fs *FileStorage) Connect() error {
	fmt.Printf("  📁 Connecting to file storage at %s\n", fs.Path)
	fs.connected = true
	return nil
}

func (fs *FileStorage) Disconnect() error {
	fmt.Println("  📁 Disconnecting from file storage")
	fs.connected = false
	return nil
}

func (fs *FileStorage) Store(key string, value interface{}) error {
	fmt.Printf("  📁 Storing to file: %s = %v\n", key, value)
	return nil
}

func (fs *FileStorage) Retrieve(key string) interface{} {
	fmt.Printf("  📁 Retrieving from file: %s\n", key)
	return map[string]interface{}{"status": "file_data", "key": key}
}

// DatabaseStorage implements DataStorage for database storage
type DatabaseStorage struct {
	Host      string
	Port      int
	Database  string
	connected bool
}

func (db *DatabaseStorage) Connect() error {
	fmt.Printf("  🗄️  Connecting to database %s at %s:%d\n", db.Database, db.Host, db.Port)
	db.connected = true
	return nil
}

func (db *DatabaseStorage) Disconnect() error {
	fmt.Println("  🗄️  Disconnecting from database")
	db.connected = false
	return nil
}

func (db *DatabaseStorage) Store(key string, value interface{}) error {
	fmt.Printf("  🗄️  Storing to database: %s = %v\n", key, value)
	return nil
}

func (db *DatabaseStorage) Retrieve(key string) interface{} {
	fmt.Printf("  🗄️  Retrieving from database: %s\n", key)
	return map[string]interface{}{"status": "db_data", "key": key}
}

// MemoryStorage implements DataStorage for in-memory storage
type MemoryStorage struct {
	Data      map[string]interface{}
	connected bool
}

func (ms *MemoryStorage) Connect() error {
	fmt.Println("  💾 Connecting to memory storage")
	ms.connected = true
	return nil
}

func (ms *MemoryStorage) Disconnect() error {
	fmt.Println("  💾 Disconnecting from memory storage")
	ms.connected = false
	return nil
}

func (ms *MemoryStorage) Store(key string, value interface{}) error {
	fmt.Printf("  💾 Storing to memory: %s = %v\n", key, value)
	ms.Data[key] = value
	return nil
}

func (ms *MemoryStorage) Retrieve(key string) interface{} {
	fmt.Printf("  💾 Retrieving from memory: %s\n", key)
	if value, exists := ms.Data[key]; exists {
		return value
	}
	return nil
}

// AdvancedPatterns demonstrates advanced Go patterns
func AdvancedPatterns() {
	fmt.Println("\n=== Advanced Patterns ===")

	// Empty interface (interface{}) - can hold any type
	fmt.Println("--- Empty Interface ---")
	var anything interface{}

	anything = 42
	fmt.Printf("Integer: %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("String: %v (type: %T)\n", anything, anything)

	anything = User{ID: 1, Username: "test"}
	fmt.Printf("Struct: %v (type: %T)\n", anything, anything)

	// Type assertion
	if user, ok := anything.(User); ok {
		fmt.Printf("Successfully asserted as User: %s\n", user.Username)
	}

	// Type switch
	fmt.Println("\n--- Type Switch ---")
	values := []interface{}{42, "hello", 3.14, true, User{ID: 1, Username: "alice"}}

	for i, value := range values {
		fmt.Printf("Value %d: ", i+1)
		switch v := value.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s (length: %d)\n", v, len(v))
		case float64:
			fmt.Printf("Float: %.2f\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		case User:
			fmt.Printf("User: %s\n", v.Username)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	// JSON marshaling/unmarshaling with struct tags
	fmt.Println("\n--- JSON Serialization ---")
	originalUser := User{
		ID:       123,
		Username: "jsonuser",
		Email:    "json@example.com",
		Role:     "user",
		Active:   true,
		Created:  time.Now(),
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(originalUser, "", "  ")
	if err != nil {
		fmt.Printf("JSON marshal error: %v\n", err)
		return
	}
	fmt.Printf("JSON representation:\n%s\n", jsonData)

	// Unmarshal from JSON
	var parsedUser User
	err = json.Unmarshal(jsonData, &parsedUser)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		return
	}
	fmt.Printf("Parsed user: %+v\n", parsedUser)
}

// CompositionExamples demonstrates composition over inheritance
func CompositionExamples() {
	fmt.Println("\n=== Composition Examples ===")

	// Creating a complex service using composition
	logger := &Logger{Level: "INFO"}
	cache := &MemoryStorage{Data: make(map[string]interface{})}

	userService := &UserService{
		Logger:  logger,
		Storage: cache,
		Config: ServiceConfig{
			MaxUsers:    1000,
			CacheExpiry: time.Hour,
		},
	}

	// Using the composed service
	fmt.Println("--- User Service Operations ---")
	user := User{
		ID:       1,
		Username: "composed_user",
		Email:    "composed@example.com",
		Role:     "user",
		Active:   true,
		Created:  time.Now(),
	}

	userService.CreateUser(user)
	retrievedUser := userService.GetUser(1)
	fmt.Printf("Retrieved user: %+v\n", retrievedUser)
}

// Logger provides logging functionality
type Logger struct {
	Level string
}

func (l *Logger) Log(level, message string) {
	fmt.Printf("[%s] %s: %s\n", time.Now().Format("15:04:05"), level, message)
}

func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

// ServiceConfig holds service configuration
type ServiceConfig struct {
	MaxUsers    int
	CacheExpiry time.Duration
}

// UserService demonstrates composition of multiple components
type UserService struct {
	Logger  *Logger       // Composed logger
	Storage DataStorage   // Composed storage (interface)
	Config  ServiceConfig // Composed configuration
}

func (us *UserService) CreateUser(user User) {
	us.Logger.Info(fmt.Sprintf("Creating user: %s", user.Username))

	// Store in cache/storage
	us.Storage.Store(fmt.Sprintf("user:%d", user.ID), user)

	us.Logger.Info(fmt.Sprintf("User created successfully: %d", user.ID))
}

func (us *UserService) GetUser(id int) *User {
	us.Logger.Info(fmt.Sprintf("Retrieving user: %d", id))

	data := us.Storage.Retrieve(fmt.Sprintf("user:%d", id))
	if data == nil {
		us.Logger.Error(fmt.Sprintf("User not found: %d", id))
		return nil
	}

	// In a real implementation, you'd properly handle type conversion
	if user, ok := data.(User); ok {
		return &user
	}

	us.Logger.Error("Invalid user data format")
	return nil
}
