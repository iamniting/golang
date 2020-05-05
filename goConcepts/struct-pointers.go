package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v
	p.X = 4
	fmt.Println(v)
	fmt.Println(*p)
	fmt.Println(p)
}
