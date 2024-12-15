package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			worker(i, ch, &wg)
		}()
	}
	close(ch)
	wg.Wait()

}

func worker(id int, task chan int, wg *sync.WaitGroup) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	defer wg.Done()
	for task := range task {
		<-ticker.C
		fmt.Printf("Worker %d processing task %d\n", id, task)
	}
}
