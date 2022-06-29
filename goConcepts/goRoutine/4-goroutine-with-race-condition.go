package main

import (
	"fmt"
	"sync"
)

// go run --race 4-goroutine-with-race-condition.go
func main() {

	wg := &sync.WaitGroup{}
	var score []int

	wg.Add(2)

	// There is a race condition as both goroutines are trying to write to the same slice without any protection
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		score = append(score, 1)
		fmt.Println("Hello 1")
	}(wg)
	// There is a race condition as both goroutines are trying to write to the same slice without any protection
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		score = append(score, 2)
		fmt.Println("Hello 2")
	}(wg)

	fmt.Println("Hello 3")
	fmt.Println("Hello 4")

	wg.Wait()
	fmt.Println(score)
}
