package main

import "fmt"

// Vertex struct
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Printf("%+v\n", v)
	v.X = 4
	fmt.Println(v.X)
}
