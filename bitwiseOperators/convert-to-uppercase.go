package main

import "fmt"

func main() {
	fmt.Printf("%c\n", 'A'&'_')
	fmt.Printf("%c\n", 'B'&'_')
	fmt.Printf("%c\n", 'a'&'_')
	fmt.Printf("%c\n", 'b'&'_')
}
