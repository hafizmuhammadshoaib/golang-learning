package main

import (
	"fmt"
	// "time"
)
func sum(value int,prevValue int,intChan chan int)  {
		fmt.Println("routine called")
		newValue := value + prevValue
		intChan <- newValue
	}

func main()  {
	intChan := make(chan int)
	prevValue := 0
	for i := 0; i <= 10; i++ {
		go sum(i,prevValue,intChan)	
		prevValue = <-intChan
	}
	fmt.Println(prevValue)
	go sum(10,10,intChan)
	// try commnent the line below
	fmt.Println(<-intChan)
	fmt.Println("function ends")
	// try uncomment the line below followed with line 22 instruction
	// time.Sleep(time.Millisecond * 1000)
}