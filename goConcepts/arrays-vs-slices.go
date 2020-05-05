package main

import "fmt"

func main() {
	// arrays always need size
	var a [10]int
	// slices dont need size at all
	var s []int

	// arrays declares with 0 at time of declarations in case of int
	fmt.Println(a)
	// slices are empty at time of declaration
	fmt.Println(s)

	fmt.Println()

	// arrays can not append as there size is fix
	// a = append(a, 5) is not permitted
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)

	// slices can append to any extend
	// s[1] = 5 is permitted if there is already a value on that index
	for i := 0; i < 15; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}
