package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func InsertNodeRecursive(root *Node, data int) *Node {
	if root == nil {
		node := &Node{data, nil, nil}
		return node
	} else if data < root.data {
		root.left = InsertNodeRecursive(root.left, data)
	} else if data > root.data {
		root.right = InsertNodeRecursive(root.right, data)
	}
	return root
}

func GetMinMaxVertical(root *Node, hd int) (int, int) {
	if root == nil {
		return hd + 1, hd - 1
	}

	var min, max int

	min, max = GetMinMaxVertical(root.left, hd-1)
	min = int(math.Min(float64(min), float64(hd)))
	max = int(math.Max(float64(max), float64(hd)))

	max, max = GetMinMaxVertical(root.right, hd+1)
	min = int(math.Min(float64(min), float64(hd)))
	max = int(math.Max(float64(max), float64(hd)))

	return min, max
}

func VerticalOrderRecursive(root *Node, hd int, vertical int) {
	if root == nil {
		return
	}

	if hd == vertical {
		fmt.Print(root.data, " ")
	}

	VerticalOrderRecursive(root.left, hd-1, vertical)
	VerticalOrderRecursive(root.right, hd+1, vertical)
}

func VerticalOrder(root *Node) {
	if root == nil {
		return
	}

	min, max := GetMinMaxVertical(root, 0)

	for ver := min; ver <= max; ver++ {
		VerticalOrderRecursive(root, 0, ver)
		fmt.Println()
	}
}

func main() {
	slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}

	var root *Node

	fmt.Println("Insert Nodes Recursively")
	for _, element := range slice {
		fmt.Print(element, " ")
		root = InsertNodeRecursive(root, element)
	}

	fmt.Println("\n\nVerticalOrder Traversal Recursively")
	VerticalOrder(root)
}
