package main

import "fmt"

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

func PostOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	PostOrderRecursive(root.left)
	PostOrderRecursive(root.right)
	fmt.Print(root.data, " ")
}

func main() {
	slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}

	var root *Node

	fmt.Println("Insert Nodes Recursively")
	for _, element := range slice {
		fmt.Print(element, " ")
		root = InsertNodeRecursive(root, element)
	}

	fmt.Println("\n\nPostOrder Traversal Recursively")
	PostOrderRecursive(root)
}
