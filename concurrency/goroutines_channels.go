// Package concurrency demonstrates Go's concurrency features
// This package covers goroutines, channels, select statements, and concurrency patterns
package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// GoroutineExamples demonstrates basic goroutine usage
func GoroutineExamples() {
	fmt.Println("=== Goroutine Examples ===")

	// Sequential execution (for comparison)
	fmt.Println("Sequential execution:")
	start := time.Now()
	task("Task 1", 1000)
	task("Task 2", 800)
	task("Task 3", 600)
	sequential := time.Since(start)
	fmt.Printf("Sequential time: %v\n", sequential)

	// Concurrent execution with goroutines
	fmt.Println("\nConcurrent execution:")
	start = time.Now()

	// Launch goroutines
	go task("Goroutine 1", 1000)
	go task("Goroutine 2", 800)
	go task("Goroutine 3", 600)

	// Wait for goroutines to complete
	time.Sleep(1200 * time.Millisecond) // Simple wait (better to use WaitGroup)
	concurrent := time.Since(start)
	fmt.Printf("Concurrent time: %v\n", concurrent)
	fmt.Printf("Speedup: %.2fx\n", float64(sequential)/float64(concurrent))

	// Anonymous goroutine
	fmt.Println("\nAnonymous goroutine:")
	go func() {
		fmt.Println("  🚀 Anonymous goroutine executed!")
	}()

	// Wait for anonymous goroutine
	time.Sleep(100 * time.Millisecond)
}

// task simulates a time-consuming operation
func task(name string, duration int) {
	fmt.Printf("  %s starting...\n", name)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	fmt.Printf("  %s completed in %dms\n", name, duration)
}

// WaitGroupExamples demonstrates proper goroutine synchronization
func WaitGroupExamples() {
	fmt.Println("\n=== WaitGroup Examples ===")

	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 5

	fmt.Printf("Starting %d workers...\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment counter

		go func(workerID int) {
			defer wg.Done() // Decrement counter when done

			// Simulate work
			workTime := rand.Intn(1000) + 500 // 500-1500ms
			fmt.Printf("  Worker %d: Starting work (%dms)\n", workerID, workTime)
			time.Sleep(time.Duration(workTime) * time.Millisecond)
			fmt.Printf("  Worker %d: Work completed\n", workerID)
		}(i) // Pass loop variable to avoid closure issues
	}

	fmt.Println("Waiting for all workers to complete...")
	wg.Wait() // Block until all goroutines call Done()
	fmt.Println("All workers completed!")
}

// ChannelExamples demonstrates channel usage for communication
func ChannelExamples() {
	fmt.Println("\n=== Channel Examples ===")

	// Unbuffered channel (synchronous)
	fmt.Println("--- Unbuffered Channel ---")
	unbuffered := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		unbuffered <- "Hello from goroutine!" // Send value
	}()

	message := <-unbuffered // Receive value (blocks until available)
	fmt.Printf("Received: %s\n", message)

	// Buffered channel (asynchronous up to buffer size)
	fmt.Println("\n--- Buffered Channel ---")
	buffered := make(chan int, 3) // Buffer size of 3

	// Send values (won't block until buffer is full)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("Sent 3 values to buffered channel")

	// Receive values
	for i := 0; i < 3; i++ {
		value := <-buffered
		fmt.Printf("Received: %d\n", value)
	}

	// Channel directions (send-only, receive-only)
	fmt.Println("\n--- Channel Directions ---")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Close channel to signal no more jobs

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

// worker demonstrates send-only and receive-only channels
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // Range over channel until closed
		fmt.Printf("  Worker %d: Processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 2                 // Send result
	}
}

// SelectExamples demonstrates the select statement for channel operations
func SelectExamples() {
	fmt.Println("\n=== Select Statement Examples ===")

	// Basic select
	fmt.Println("--- Basic Select ---")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(800 * time.Millisecond)
		c2 <- "Message from channel 2"
	}()

	// Select waits for the first available channel
	select {
	case msg1 := <-c1:
		fmt.Printf("Received: %s\n", msg1)
	case msg2 := <-c2:
		fmt.Printf("Received: %s\n", msg2)
	}

	// Select with timeout
	fmt.Println("\n--- Select with Timeout ---")
	slowChannel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Too slow!
		slowChannel <- "Finally arrived"
	}()

	select {
	case msg := <-slowChannel:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Operation took too long")
	}

	// Non-blocking select with default
	fmt.Println("\n--- Non-blocking Select ---")
	nonBlocking := make(chan int)

	select {
	case value := <-nonBlocking:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("No value available, continuing...")
	}
}

// ProducerConsumerPattern demonstrates a common concurrency pattern
func ProducerConsumerPattern() {
	fmt.Println("\n=== Producer-Consumer Pattern ===")

	// Create channels
	jobs := make(chan Job, 10)       // Buffered channel for jobs
	results := make(chan Result, 10) // Buffered channel for results

	// Start consumers (workers)
	numWorkers := 3
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go consumer(i, jobs, results, &wg)
	}

	// Start producer
	go producer(jobs, 10)

	// Start result collector
	go resultCollector(results, 10)

	// Wait for consumers to finish
	wg.Wait()
	close(results)

	fmt.Println("All jobs processed!")
}

// Job represents work to be done
type Job struct {
	ID   int
	Data string
}

// Result represents the result of processed work
type Result struct {
	JobID    int
	Output   string
	Duration time.Duration
}

// producer generates jobs and sends them to the jobs channel
func producer(jobs chan<- Job, numJobs int) {
	defer close(jobs)

	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("data-%d", i),
		}
		fmt.Printf("📤 Producing job %d\n", job.ID)
		jobs <- job
		time.Sleep(100 * time.Millisecond) // Simulate production time
	}
}

// consumer processes jobs from the jobs channel
func consumer(workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		start := time.Now()
		fmt.Printf("  👷 Worker %d: Processing job %d\n", workerID, job.ID)

		// Simulate processing time
		processingTime := time.Duration(rand.Intn(500)+200) * time.Millisecond
		time.Sleep(processingTime)

		result := Result{
			JobID:    job.ID,
			Output:   fmt.Sprintf("processed-%s", job.Data),
			Duration: time.Since(start),
		}

		results <- result
		fmt.Printf("  ✅ Worker %d: Completed job %d in %v\n",
			workerID, job.ID, result.Duration)
	}
}

// resultCollector collects and processes results
func resultCollector(results <-chan Result, expectedResults int) {
	collected := 0
	totalDuration := time.Duration(0)

	for result := range results {
		collected++
		totalDuration += result.Duration
		fmt.Printf("📥 Result %d: %s (took %v)\n",
			result.JobID, result.Output, result.Duration)

		if collected >= expectedResults {
			break
		}
	}

	avgDuration := totalDuration / time.Duration(collected)
	fmt.Printf("📊 Collected %d results, average duration: %v\n",
		collected, avgDuration)
}

// ContextExamples demonstrates context usage for cancellation and timeouts
func ContextExamples() {
	fmt.Println("\n=== Context Examples ===")

	// Context with timeout
	fmt.Println("--- Context with Timeout ---")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go longRunningTask(ctx, "Task 1", 1*time.Second) // Should complete
	go longRunningTask(ctx, "Task 2", 3*time.Second) // Should be cancelled

	time.Sleep(3 * time.Second) // Wait for demonstration

	// Context with cancellation
	fmt.Println("\n--- Context with Cancellation ---")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go longRunningTask(ctx2, "Task 3", 5*time.Second)

	// Cancel after 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Cancelling context...")
	cancel2()

	time.Sleep(500 * time.Millisecond) // Wait for cancellation to take effect
}

// longRunningTask simulates a task that can be cancelled via context
func longRunningTask(ctx context.Context, name string, duration time.Duration) {
	fmt.Printf("  🏃 %s: Starting (expected duration: %v)\n", name, duration)

	select {
	case <-time.After(duration):
		fmt.Printf("  ✅ %s: Completed successfully\n", name)
	case <-ctx.Done():
		fmt.Printf("  ❌ %s: Cancelled (%v)\n", name, ctx.Err())
	}
}

// PipelinePattern demonstrates a concurrent pipeline
func PipelinePattern() {
	fmt.Println("\n=== Pipeline Pattern ===")

	// Create pipeline stages
	numbers := generateNumbers(1, 10)
	squared := squareNumbers(numbers)
	filtered := filterEven(squared)

	// Consume final results
	fmt.Println("Pipeline results (even squares):")
	for result := range filtered {
		fmt.Printf("  %d\n", result)
	}
}

// generateNumbers creates a channel of numbers (stage 1)
func generateNumbers(start, end int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i <= end; i++ {
			fmt.Printf("  📤 Generating: %d\n", i)
			out <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

// squareNumbers squares input numbers (stage 2)
func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			squared := n * n
			fmt.Printf("  🔢 Squaring %d = %d\n", n, squared)
			out <- squared
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return out
}

// filterEven filters even numbers (stage 3)
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				fmt.Printf("  ✅ Filtering (even): %d\n", n)
				out <- n
			} else {
				fmt.Printf("  ❌ Filtering (odd): %d\n", n)
			}
			time.Sleep(30 * time.Millisecond)
		}
	}()
	return out
}

// WorkerPoolPattern demonstrates a worker pool for processing tasks
func WorkerPoolPattern() {
	fmt.Println("\n=== Worker Pool Pattern ===")

	const numWorkers = 3
	const numJobs = 8

	// Create channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	fmt.Printf("Starting %d workers...\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go poolWorker(w, jobs, results)
	}

	// Send jobs
	fmt.Printf("Sending %d jobs...\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	fmt.Println("Collecting results:")
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("  Result: %d\n", result)
	}
}

// poolWorker processes jobs from the worker pool
func poolWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("  👷 Worker %d: Processing job %d\n", id, job)

		// Simulate work
		time.Sleep(time.Duration(rand.Intn(300)+200) * time.Millisecond)

		// Return result (fibonacci number for demonstration)
		result := fibonacci(job)
		results <- result

		fmt.Printf("  ✅ Worker %d: Job %d completed (result: %d)\n", id, job, result)
	}
}

// fibonacci calculates fibonacci number (simple recursive implementation)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
