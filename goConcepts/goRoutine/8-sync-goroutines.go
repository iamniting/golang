package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut1 := &sync.Mutex{}
	mut2 := &sync.Mutex{}
	mut3 := &sync.Mutex{}

	wg.Add(9)

	// lock routine B and C
	mut2.Lock()
	mut3.Lock()

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			mut1.Lock()
			fmt.Println("A")
			mut2.Unlock()
		}()
		go func() {
			defer wg.Done()
			mut2.Lock()
			fmt.Println("B")
			mut3.Unlock()
		}()
		go func() {
			defer wg.Done()
			mut3.Lock()
			fmt.Println("C")
			mut1.Unlock()
		}()
	}

	wg.Wait()
}
