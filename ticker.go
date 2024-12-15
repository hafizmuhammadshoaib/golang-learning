package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5"}

	for _, task := range tasks {
		<-ticker.C // Wait for the next tick
		fmt.Printf("Processing %s\n", task)
	}

}
