package main

import (
	"fmt"
	"sort"
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

func VerticalOrderIterative(root *Node) {
	if root == nil {
		return
	}

	type vertical struct {
		node *Node
		hd   int
	}

	m := make(map[int][]int)

	var queue []vertical
	queue = append(queue, vertical{root, 0})

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		// push hd as key and node.data as list of value in map
		m[parent.hd] = append(m[parent.hd], parent.node.data)

		if parent.node.left != nil {
			queue = append(queue, vertical{(parent.node).left, parent.hd - 1})
		}
		if parent.node.right != nil {
			queue = append(queue, vertical{(parent.node).right, parent.hd + 1})
		}
	}

	// get and sort the keys of map
	var hds []int
	for hd := range m {
		hds = append(hds, hd)
	}
	sort.Ints(hds)

	// print vertical order
	for _, hd := range hds {
		fmt.Println(m[hd])
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

	fmt.Println("\n\nVerticalOrder Traversal Iteratively")
	VerticalOrderIterative(root)
}
