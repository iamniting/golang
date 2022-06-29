package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello 1")
	go fmt.Println("Hello 2")
	fmt.Println("Hello 3")
	fmt.Println("Hello 4")
	// it does not make sure that goroutine is complete
	time.Sleep(1 * time.Second)
}
