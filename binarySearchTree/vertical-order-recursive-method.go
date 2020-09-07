package main

import (
	"fmt"
	"math"
)

// Node defines a node of bst
type Node struct {
	data  int
	left  *Node
	right *Node
}

// InsertNodeRecursive insert a node in tree
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

// GetMinMaxVertical return min and max vertical in tree
func GetMinMaxVertical(root *Node, hd int) (int, int) {
	if root == nil {
		return hd + 1, hd - 1
	}

	lmin, lmax := GetMinMaxVertical(root.left, hd-1)
	lmin = int(math.Min(float64(lmin), float64(hd)))
	lmax = int(math.Max(float64(lmax), float64(hd)))

	rmin, rmax := GetMinMaxVertical(root.right, hd+1)
	rmin = int(math.Min(float64(rmin), float64(hd)))
	rmax = int(math.Max(float64(rmax), float64(hd)))

	min := int(math.Min(float64(lmin), float64(rmin)))
	max := int(math.Max(float64(lmax), float64(rmax)))

	return min, max
}

// VerticalOrderRecursive iterate tree in vertical order
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

// VerticalOrder iterate tree in vertical order
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
