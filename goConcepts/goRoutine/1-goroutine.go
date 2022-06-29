package main

import "fmt"

func main() {
	go fmt.Println("Hello 1")
	go fmt.Println("Hello 2")
	fmt.Println("Hello 3")
	fmt.Println("Hello 4")
}
