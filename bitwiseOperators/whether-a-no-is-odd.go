package main

import "fmt"

func isNumberOdd(n int) bool {
	// last bit will be always set for odd no
	return (n&1 == 1)
}

func main() {
	for i := 0; i <= 16; i++ {
		fmt.Println(i, isNumberOdd(i))
	}
}
