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

func LevelOrderIterative(root *Node) {
	if root == nil {
		return
	}
	var queue []*Node
	queue = append(queue, root)

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		if parent.left == nil && parent.right == nil {
			fmt.Print(parent.data, " ")
		}

		if parent.left != nil {
			queue = append(queue, parent.left)
		}
		if parent.right != nil {
			queue = append(queue, parent.right)
		}
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

	fmt.Println("\n\nLeaf Nodes Iteratively")
	LevelOrderIterative(root)
}
