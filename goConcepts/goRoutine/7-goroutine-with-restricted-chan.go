package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Read only chan, can not close the chan here
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// Reading from chan
		fmt.Println(<-myCh)
		fmt.Println("Hello 1")
		wg.Done()
	}(myCh, wg)

	// Send only chan, can close the chan here
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// Sending to chan
		myCh <- 5
		close(myCh)
		fmt.Println("Hello 2")
		wg.Done()
	}(myCh, wg)

	fmt.Println("Hello 3")
	fmt.Println("Hello 4")

	wg.Wait()
}
