package main

import "fmt"

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

// InOrderRecursive traverse tree inorder fashion
func InOrderRecursive(root *Node) {
	if root == nil {
		return
	}
	InOrderRecursive(root.left)
	fmt.Print(root.data, " ")
	InOrderRecursive(root.right)
}

// DeleteNodeIterative delete a node in tree
func DeleteNodeIterative(root *Node, data int) *Node {
	node := root
	parent := root

	// find element node and parent
	for node != nil {
		if node.data == data {
			break
		} else if node.data > data {
			parent = node
			node = node.left
		} else if node.data < data {
			parent = node
			node = node.right
		}
	}

	if node == nil {
		return root
	}

	// if leaf node
	if node.left == nil && node.right == nil {
		if parent.left == node {
			parent.left = nil
		} else if parent.right == node {
			parent.right = nil
		} else if parent == node && root == node {
			return nil
		}
		// if only left child
	} else if node.left != nil && node.right == nil {
		if parent.left == node {
			parent.left = node.left
			node.left = nil
		} else if parent.right == node {
			parent.right = node.left
			node.left = nil
		} else if parent == node && root == node {
			node = node.left
			return node
		}
		// if only right child
	} else if node.left == nil && node.right != nil {
		if parent.left == node {
			parent.left = node.right
			node.right = nil
		} else if parent.right == node {
			parent.right = node.right
			node.right = nil
		} else if parent == node && root == node {
			node = node.right
			return node
		}
		// if both childs are present
	} else if node.left != nil && node.right != nil {
		// get right most child of left sub tree
		var successor *Node = node.left
		var succParent *Node = node

		for successor.right != nil {
			succParent = successor
			successor = successor.right
		}

		if succParent == node {
			node.data = successor.data
			node.left = successor.left
			successor = nil
		} else if succParent != node {
			succParent.right = successor.left
			successor.left = nil
			node.data = successor.data
		}
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

	fmt.Printf("\n\nDelete Nodes Iteratively")
	del := []int{5, 100, 90, 30}

	for _, element := range del {
		fmt.Printf("\n\nDelete %v", element)
		root = DeleteNodeIterative(root, element)
		fmt.Println("\n\nInOrder Traversal Recursively After Deletion")
		InOrderRecursive(root)
	}
}
