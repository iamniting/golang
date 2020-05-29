package main

import "fmt"

func divideByTwo(num int) int {
	/*
	   64 = 1000000
	   32 = 0100000
	   64 >> 1, which is 32
	*/
	return num >> 1
}

func main() {
	for i := 0; i <= 16; i++ {
		fmt.Println(i, divideByTwo(i))
	}
}
