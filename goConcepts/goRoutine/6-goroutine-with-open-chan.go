package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Read and send chan, can close the chan here
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		close(myCh)
		fmt.Println("Hello 1")
		wg.Done()
	}(myCh, wg)

	// Read and send chan, can close the chan here
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		fmt.Println("Hello 2")
		wg.Done()
	}(myCh, wg)

	fmt.Println("Hello 3")
	fmt.Println("Hello 4")

	wg.Wait()
}
