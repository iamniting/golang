package main

import "fmt"

// Node defines a node of bst
type Node struct {
	data  int
	left  *Node
	right *Node
}

// InsertNodeIterative insert a node in tree
func InsertNodeIterative(root *Node, data int) *Node {
	if root == nil {
		node := &Node{data, nil, nil}
		return node
	}

	node := root

	for node != nil {
		if node.data > data && node.left == nil {
			node.left = &Node{data, nil, nil}
			break
		} else if node.data < data && node.right == nil {
			node.right = &Node{data, nil, nil}
			break
		} else if node.data > data {
			node = node.left
		} else if node.data < data {
			node = node.right
		}
	}
	return root
}

// InOrderRecursive traverse tree inorder fashion
func InOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	InOrderRecursive(root.left)
	fmt.Print(root.data, " ")
	InOrderRecursive(root.right)
}

func main() {
	slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}

	var root *Node

	fmt.Println("Insert Nodes Iteratively")
	for _, element := range slice {
		fmt.Print(element, " ")
		root = InsertNodeIterative(root, element)
	}

	fmt.Println("\n\nInOrder Traversal Recursively")
	InOrderRecursive(root)
}
