# Concurrency: 
## What is Concurrency?

Concurrency is about **dealing with lots of things at once**. It's the composition of independently executing processes/tasks. In programming, it means structuring your program so that multiple tasks can be in progress simultaneously, even if they're not necessarily running at the exact same time.

### Real-World Analogy: Coffee Shop
Imagine a coffee shop with one barista:
- **Sequential**: Take order → Make coffee → Serve → Take next order
- **Concurrent**: Take order → Start coffee machine → While coffee brews, take another order → Serve first coffee → Continue with second order

The barista isn't doing everything simultaneously, but they're managing multiple tasks efficiently.

### Go's Concurrency Model

Go uses **goroutines** and **channels** for concurrency:

1. **Goroutines**: Lightweight threads managed by Go runtime
2. **Channels**: Communication mechanism between goroutines
3. **Select Statement**: Multiplexing for channels

### Key Concepts:

#### 1. Goroutines
```go
// Sequential execution
func sequential() {
    task1()
    task2()
    task3()
}

// Concurrent execution
func concurrent() {
    go task1()  // Runs in separate goroutine
    go task2()  // Runs in separate goroutine
    task3()     // Runs in main goroutine
}
```

#### 2. Channels
```go
// Creating channels
ch := make(chan int)        // Unbuffered channel
buffered := make(chan int, 3) // Buffered channel

// Sending and receiving
ch <- 42        // Send value to channel
value := <-ch   // Receive value from channel
```

#### 3. Select Statement
```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
}
```

### Practical Example: Web Scraper
```go
package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

func fetchURL(url string, wg *sync.WaitGroup, results chan<- string) {
    defer wg.Done()
    
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        results <- fmt.Sprintf("%s: Error - %v", url, err)
        return
    }
    defer resp.Body.Close()
    
    duration := time.Since(start)
    results <- fmt.Sprintf("%s: %s (took %v)", url, resp.Status, duration)
}

func main() {
    urls := []string{
        "https://golang.org",
        "https://github.com",
        "https://stackoverflow.com",
    }
    
    var wg sync.WaitGroup
    results := make(chan string, len(urls))
    
    // Launch goroutines
    for _, url := range urls {
        wg.Add(1)
        go fetchURL(url, &wg, results)
    }
    
    // Wait for all goroutines to complete
    wg.Wait()
    close(results)
    
    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}
```

# Parallelism:
## What is Parallelism?

Parallelism is about **doing lots of things at once**. It's the simultaneous execution of multiple tasks using multiple CPU cores or processors.

### Key Difference: Concurrency vs Parallelism

**Analogy: Highway Traffic**
- **Concurrency**: One lane with cars taking turns efficiently (time-slicing)
- **Parallelism**: Multiple lanes with cars moving simultaneously

```
Concurrency (Single Core):     Parallelism (Multi-Core):
Task A: --|-----|-----         Task A: ----------
Task B:   --|-----|---         Task B: ----------
Task C:     --|-----|-         Task C: ----------
Time:   ================       Time:   ==========
```

### When Go Achieves Parallelism

Go can achieve true parallelism when:
1. Multiple CPU cores are available
2. `GOMAXPROCS` > 1 (default: number of CPU cores)
3. Goroutines are CPU-intensive (not just I/O bound)

### Example: CPU-Intensive Parallel Work
```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func cpuIntensiveWork(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    // Simulate CPU-intensive work
    start := time.Now()
    sum := 0
    for i := 0; i < 1000000000; i++ {
        sum += i
    }
    
    fmt.Printf("Worker %d finished in %v (sum: %d)\n", 
               id, time.Since(start), sum)
}

func main() {
    fmt.Printf("CPU cores: %d\n", runtime.NumCPU())
    fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
    
    var wg sync.WaitGroup
    numWorkers := runtime.NumCPU()
    
    start := time.Now()
    
    // Launch parallel workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go cpuIntensiveWork(i, &wg)
    }
    
    wg.Wait()
    fmt.Printf("Total time: %v\n", time.Since(start))
}
```

### Concurrency Patterns in Go

#### 1. Worker Pool Pattern
```go
func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}
```

#### 2. Fan-out/Fan-in Pattern
```go
func fanOut(input <-chan int) (<-chan int, <-chan int) {
    out1 := make(chan int)
    out2 := make(chan int)
    
    go func() {
        defer close(out1)
        defer close(out2)
        for val := range input {
            out1 <- val
            out2 <- val
        }
    }()
    
    return out1, out2
}

func fanIn(input1, input2 <-chan int) <-chan int {
    output := make(chan int)
    
    go func() {
        defer close(output)
        for {
            select {
            case val, ok := <-input1:
                if !ok {
                    input1 = nil
                } else {
                    output <- val
                }
            case val, ok := <-input2:
                if !ok {
                    input2 = nil
                } else {
                    output <- val
                }
            }
            if input1 == nil && input2 == nil {
                break
            }
        }
    }()
    
    return output
}
```

### Best Practices

1. **Don't communicate by sharing memory; share memory by communicating**
2. **Use channels for communication, mutexes for protecting state**
3. **Prefer goroutines over threads**
4. **Use context for cancellation and timeouts**
5. **Always handle channel closure**
6. **Avoid goroutine leaks**

### Common Pitfalls

1. **Race Conditions**: Multiple goroutines accessing shared data
2. **Deadlocks**: Goroutines waiting for each other indefinitely
3. **Goroutine Leaks**: Goroutines that never terminate
4. **Channel Blocking**: Sending to full channel or receiving from empty channel

### Summary

- **Concurrency**: Design pattern for managing multiple tasks
- **Parallelism**: Actual simultaneous execution
- **Go provides**: Goroutines, channels, and select for elegant concurrency
- **Key insight**: Concurrency enables parallelism, but they're different concepts

## Buffered vs Unbuffered Channels

### Unbuffered Channels (Synchronous)

An **unbuffered channel** has no capacity to store values. It requires both a sender and receiver to be ready at the same time.

**Analogy: Phone Call**
- When you make a phone call, both parties must be available simultaneously
- The call happens in real-time - no storage/buffering
- If the receiver isn't available, the call fails

```go
package main

import (
    "fmt"
    "time"
)

func unbufferedExample() {
    ch := make(chan string) // Unbuffered channel
    
    // This would cause a deadlock if uncommented:
    // ch <- "Hello" // Blocks forever - no receiver ready
    
    // Proper usage with goroutine
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Hello from goroutine" // Sends when receiver is ready
    }()
    
    fmt.Println("Waiting for message...")
    msg := <-ch // Blocks until sender sends
    fmt.Println("Received:", msg)
}
```

### Buffered Channels (Asynchronous)

A **buffered channel** has a capacity to store values. Senders can send values without an immediate receiver (until buffer is full).

**Analogy: Mailbox**
- You can drop letters in a mailbox even if the recipient isn't home
- The mailbox has limited capacity
- If mailbox is full, you must wait for space

```go
package main

import (
    "fmt"
    "time"
)

func bufferedExample() {
    ch := make(chan string, 3) // Buffered channel with capacity 3
    
    // Can send up to 3 values without blocking
    ch <- "Message 1"
    ch <- "Message 2"
    ch <- "Message 3"
    
    fmt.Printf("Buffer length: %d, capacity: %d\n", len(ch), cap(ch))
    
    // This would block because buffer is full:
    // ch <- "Message 4"
    
    // Receive messages
    for i := 0; i < 3; i++ {
        msg := <-ch
        fmt.Println("Received:", msg)
    }
}
```

### Detailed Comparison

#### Unbuffered Channel Behavior
```go
func unbufferedBehavior() {
    ch := make(chan int)
    
    // Sender blocks until receiver is ready
    go func() {
        fmt.Println("Sending 42...")
        ch <- 42 // Blocks here until main goroutine receives
        fmt.Println("Sent 42!")
    }()
    
    time.Sleep(2 * time.Second) // Simulate work
    fmt.Println("Ready to receive...")
    value := <-ch // Unblocks the sender
    fmt.Printf("Received: %d\n", value)
}
```

#### Buffered Channel Behavior
```go
func bufferedBehavior() {
    ch := make(chan int, 2) // Buffer size 2
    
    go func() {
        for i := 1; i <= 4; i++ {
            fmt.Printf("Sending %d...\n", i)
            ch <- i
            fmt.Printf("Sent %d! Buffer length: %d\n", i, len(ch))
            
            if i == 2 {
                fmt.Println("Buffer full! Next send will block...")
            }
        }
    }()
    
    time.Sleep(3 * time.Second) // Let first 2 sends complete
    
    // Receive values
    for i := 0; i < 4; i++ {
        value := <-ch
        fmt.Printf("Received: %d\n", value)
        time.Sleep(1 * time.Second)
    }
}
```

### Practical Examples

#### Example 1: Producer-Consumer with Buffered Channel
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func producer(ch chan<- int, id int) {
    for i := 0; i < 5; i++ {
        value := rand.Intn(100)
        fmt.Printf("Producer %d: producing %d\n", id, value)
        ch <- value
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    }
}

func consumer(ch <-chan int, id int) {
    for value := range ch {
        fmt.Printf("Consumer %d: consumed %d\n", id, value)
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
}

func producerConsumerExample() {
    // Buffered channel allows producers to work ahead of consumers
    ch := make(chan int, 10)
    
    // Start producers
    go producer(ch, 1)
    go producer(ch, 2)
    
    // Start consumer
    go consumer(ch, 1)
    
    time.Sleep(5 * time.Second)
    close(ch)
    time.Sleep(1 * time.Second) // Let consumer finish
}
```

#### Example 2: Rate Limiting with Buffered Channel
```go
func rateLimiter() {
    // Create a buffered channel as a token bucket
    tokens := make(chan struct{}, 3) // Allow 3 concurrent operations
    
    // Fill the bucket with tokens
    for i := 0; i < 3; i++ {
        tokens <- struct{}{}
    }
    
    // Simulate multiple requests
    for i := 0; i < 10; i++ {
        go func(id int) {
            // Acquire token (blocks if none available)
            <-tokens
            fmt.Printf("Processing request %d\n", id)
            
            // Simulate work
            time.Sleep(2 * time.Second)
            
            // Return token
            tokens <- struct{}{}
            fmt.Printf("Finished request %d\n", id)
        }(i)
    }
    
    time.Sleep(15 * time.Second)
}
```

### When to Use Each Type

#### Use Unbuffered Channels When:
- You need synchronization between goroutines
- You want to ensure immediate delivery
- You need back-pressure (slow receiver slows down sender)
- Implementing request-response patterns

```go
func requestResponse() {
    request := make(chan string)
    response := make(chan string)
    
    // Server goroutine
    go func() {
        req := <-request // Wait for request
        fmt.Printf("Processing: %s\n", req)
        response <- "Processed: " + req // Send response
    }()
    
    // Client
    request <- "Hello"
    resp := <-response
    fmt.Println("Got response:", resp)
}
```

#### Use Buffered Channels When:
- You want to decouple sender and receiver timing
- You need to handle bursts of data
- Implementing worker pools or rate limiting
- You know the expected capacity

```go
func burstyWork() {
    work := make(chan string, 100) // Handle bursts
    
    // Worker
    go func() {
        for job := range work {
            fmt.Printf("Processing: %s\n", job)
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    // Burst of work
    for i := 0; i < 50; i++ {
        work <- fmt.Sprintf("Job %d", i)
    }
    
    close(work)
    time.Sleep(6 * time.Second)
}
```

### Key Differences Summary

| Aspect | Unbuffered | Buffered |
|--------|------------|----------|
| **Synchronization** | Synchronous (blocks until both ready) | Asynchronous (up to buffer size) |
| **Capacity** | 0 | Specified size (> 0) |
| **Send Blocking** | Always blocks until receiver ready | Blocks only when buffer full |
| **Receive Blocking** | Always blocks until sender ready | Blocks only when buffer empty |
| **Use Case** | Synchronization, immediate delivery | Decoupling, handling bursts |
| **Memory Usage** | Minimal | Uses memory for buffer |

### Common Patterns and Gotchas

#### Gotcha 1: Deadlock with Unbuffered Channel
```go
// WRONG - This will deadlock
func deadlockExample() {
    ch := make(chan int)
    ch <- 42 // Blocks forever - no receiver
    fmt.Println(<-ch)
}

// CORRECT - Use goroutine or buffered channel
func fixedExample() {
    ch := make(chan int, 1) // Buffered
    ch <- 42
    fmt.Println(<-ch)
}
```

#### Gotcha 2: Forgetting to Close Channels
```go
func properChannelUsage() {
    ch := make(chan int, 5)
    
    go func() {
        defer close(ch) // Always close when done sending
        for i := 0; i < 5; i++ {
            ch <- i
        }
    }()
    
    // Range automatically stops when channel closes
    for value := range ch {
        fmt.Println(value)
    }
}
```