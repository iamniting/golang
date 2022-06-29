package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Hello 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Hello 2")
	}()

	fmt.Println("Hello 3")
	fmt.Println("Hello 4")

	// It makes sure goroutine is over
	wg.Wait()
}
