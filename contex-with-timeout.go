package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go worker(ctx)
	<-ctx.Done()
	if ctx.Err() != nil {
		fmt.Println("Task was canceled:", ctx.Err())
	} else {
		fmt.Println("Task completed successfully.")
	}
}
func worker(ctx context.Context) {
	select {
	case <-time.After(12 * time.Second): // Simulate long task
		fmt.Println("Worker completed the task")
	case <-ctx.Done(): // Context is canceled or times out
		fmt.Println("Worker received cancel signal")
	}
}
