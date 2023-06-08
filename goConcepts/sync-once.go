package main

import (
	"fmt"
	"sync"
)

func main() {

	var once sync.Once

	printHiHello(&once)
	printHiHello(&once)
	printHiHello(&once)
}

func printHiHello(once *sync.Once) {

	// The "Hi" is printed only once, and "Hello" is printed for each call to printHiHello
	once.Do(func() {
		fmt.Println("Hi")
	})
	fmt.Println("Hello")
}
