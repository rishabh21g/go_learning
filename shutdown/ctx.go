package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// HERE IS THE EXAMPLE OF GRACEFUL SHUTDOWN USING CONTEXT IN GO

func ctxShutdown() {

	server := http.Server{
		Addr: ":3000",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// stimulating the server handling the long req
		time.Sleep(4 * time.Second)
		fmt.Fprintln(w, "Hellp from the Go Server")
	})
	// running the server in the goroutine
	go func() {
		fmt.Print("Server is running on the port : 3000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Listen and Serve : %v\n", err)
		}
	}()

	// wait for a signal interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// create a ctx with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shoutdown failed: %v\n", err)
	}
	fmt.Println("GRACEFULLY SHUTDOWN")
}

func main() {
	ctxShutdown()
}
