package main

import (
	"fmt"
	"sync"
)

// go run --race 5-goroutine-without-race-condition-using-mutex.go
func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score []int

	wg.Add(2)

	// There is no race condition as both goroutines are trying to write to the same slice with mutex protection
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		fmt.Println("Hello 1")
	}(wg, mut)
	// There is no race condition as both goroutines are trying to write to the same slice with mutex protection
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		fmt.Println("Hello 2")
	}(wg, mut)

	fmt.Println("Hello 3")
	fmt.Println("Hello 4")

	wg.Wait()
	fmt.Println(score)
}
