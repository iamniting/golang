package main

import "fmt"

func main() {
	// explicit declaration
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	fmt.Println()

	// implicit declaration
	slice := []int{1, 2, 3, 4}
	printSlice(slice)

	// Add more elements
	slice = append(slice, 5, 6)
	printSlice(slice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
