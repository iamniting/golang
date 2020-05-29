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

func InOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	InOrderRecursive(root.left)
	fmt.Print(root.data, " ")
	InOrderRecursive(root.right)
}

func DeleteNodeRecursive(root *Node, data int) *Node {
	if root == nil {
		return root
	} else if root.data > data {
		root.left = DeleteNodeRecursive(root.left, data)
	} else if root.data < data {
		root.right = DeleteNodeRecursive(root.right, data)
		// if leaf node which need to be deleted
	} else if root.data == data && root.left == nil && root.right == nil {
		return nil
		// if node to be deleted having one or both child
	} else {
		// node with only left child
		if root.left == nil {
			return root.right
			// node with only right child
		} else if root.right == nil {
			return root.left
		}

		// get max value of left sub tree
		max := root.left
		for max.right != nil {
			max = max.right
		}

		// replace root data with the max value
		root.data = max.data

		// delete max value of left sub tree
		root.left = DeleteNodeRecursive(root.left, max.data)
	}
	return root
}

func main() {
	slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}

	var root *Node

	fmt.Println("Insert Nodes Recursively")
	for _, element := range slice {
		fmt.Print(element, " ")
		root = InsertNodeRecursive(root, element)
	}

	fmt.Printf("\n\nDelete Nodes Recursively")
	del := []int{5, 100, 90, 30}

	for _, element := range del {
		fmt.Printf("\n\nDelete %v", element)
		root = DeleteNodeRecursive(root, element)
		fmt.Println("\n\nInOrder Traversal Recursively After Deletion")
		InOrderRecursive(root)
	}
}
