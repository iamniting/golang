package main

import "fmt"

func main() {
	// A defer statement defers the execution of a function
	// until the surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("hello")
}
