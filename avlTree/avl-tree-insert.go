package main

import (
	"fmt"
	"math"
)

type Node struct {
	data int
	// balanced factor
	bf    int
	left  *Node
	right *Node
}

func InsertInAvl(root *Node, data int) *Node {
	// insert element like bst
	if root == nil {
		root = &Node{data, 0, nil, nil}
	} else if root.data > data {
		root.left = InsertInAvl(root.left, data)
	} else if root.data < data {
		root.right = InsertInAvl(root.right, data)
	}

	// calculate balance factor of node
	root.bf = GetBf(root)

	// already balanced
	if root.bf == 0 || root.bf == -1 || root.bf == 1 {
		return root
	}

	// left left imbalanced
	if root.bf > 1 && root.left.bf > 0 {
		return RightRotation(root)
	}

	// right right imbalanced
	if root.bf < -1 && root.right.bf < 0 {
		return LeftRotation(root)
	}

	// left right imbalanced
	if root.bf > 1 && root.left.bf < 0 {
		root.left = LeftRotation(root.left)
		return RightRotation(root)
	}

	// right left imbalanced
	if root.bf < -1 && root.right.bf > 0 {
		root.right = RightRotation(root.right)
		return LeftRotation(root)
	}

	return root
}

func LeftRotation(root *Node) *Node {
	x := root.right
	y := x.left

	// rotate the root to left side
	x.left = root
	root.right = y

	// calculate the balanced factor of root and x
	root.bf = GetBf(root)
	x.bf = GetBf(x)

	return x
}

func RightRotation(root *Node) *Node {
	x := root.left
	y := x.right

	// rotate the root to right side
	x.right = root
	root.left = y

	// calculate the balanced factor of root and x
	root.bf = GetBf(root)
	x.bf = GetBf(x)

	return x
}

func GetBf(root *Node) int {
	hleft := HeightRecursive(root.left)
	hright := HeightRecursive(root.right)
	root.bf = hleft - hright

	return root.bf
}

func HeightRecursive(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(
		float64(HeightRecursive(root.left)),
		float64(HeightRecursive(root.right))))
}

func LevelOrderRecursive(root *Node, level int) {
	if level == 1 && root != nil {
		fmt.Printf("(%v,%v) ", root.data, root.bf)
		return
	}

	if level == 1 && root == nil {
		fmt.Print("(no,no)", " ")
		return
	}

	if root != nil && root.left != nil {
		LevelOrderRecursive(root.left, level-1)
	} else {
		LevelOrderRecursive(nil, level-1)
	}
	if root != nil && root.right != nil {
		LevelOrderRecursive(root.right, level-1)
	} else {
		LevelOrderRecursive(nil, level-1)
	}
}

func LevelOrder(root *Node) {
	height := HeightRecursive(root)

	for h := 1; h <= height; h++ {
		LevelOrderRecursive(root, h)
		fmt.Println()
	}
}

func main() {
	slice := []int{20, 30, 40, 35, 38, 37, 10}
	var root *Node

	fmt.Println("Insert Nodes in AVL Tree Recursively")
	for _, element := range slice {
		fmt.Print(element, " ")
		root = InsertInAvl(root, element)
	}

	fmt.Println("\n\nLevelOrder Traversal Recursively")
	LevelOrder(root)
}
