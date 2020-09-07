package avl

import (
	"fmt"
	"math"
)

// Node defines avl tree struct
type Node struct {
	Val int
	// balanced factor
	Bf    int
	Left  *Node
	Right *Node
}

// InsertInAvl insert a node in avl tree
func InsertInAvl(root *Node, Val int) *Node {
	// insert element like bst
	if root == nil {
		root = &Node{Val, 0, nil, nil}
	} else if root.Val > Val {
		root.Left = InsertInAvl(root.Left, Val)
	} else if root.Val < Val {
		root.Right = InsertInAvl(root.Right, Val)
	}

	// calculate balanced factor of node
	root.Bf = GetBalancedFactor(root)

	// already balanced
	if root.Bf == 0 || root.Bf == -1 || root.Bf == 1 {
		return root
	}

	// left left imbalanced
	if root.Bf > 1 && root.Left.Bf > 0 {
		return RightRotation(root)
	}

	// right right imbalanced
	if root.Bf < -1 && root.Right.Bf < 0 {
		return LeftRotation(root)
	}

	// left right imbalanced
	if root.Bf > 1 && root.Left.Bf < 0 {
		root.Left = LeftRotation(root.Left)
		return RightRotation(root)
	}

	// right left imbalanced
	if root.Bf < -1 && root.Right.Bf > 0 {
		root.Right = RightRotation(root.Right)
		return LeftRotation(root)
	}

	return root
}

// DeleteInAvl delete a node from avl tree
func DeleteInAvl(root *Node, Val int) *Node {
	if root == nil {
		return root
	} else if root.Val > Val {
		root.Left = DeleteInAvl(root.Left, Val)
	} else if root.Val < Val {
		root.Right = DeleteInAvl(root.Right, Val)
		// if leaf node which need to be deleted
	} else if root.Val == Val && root.Left == nil && root.Right == nil {
		return nil
		// if node to be deleted having one or both child
	} else {
		// node with only right child
		if root.Left == nil {
			root = root.Right
			// node with only left child
		} else if root.Right == nil {
			root = root.Left
		}

		// get max value of left sub tree
		max := root.Left
		for max.Right != nil {
			max = max.Right
		}

		// replace root val with the max value
		root.Val = max.Val

		// delete max value of left sub tree
		root.Left = DeleteInAvl(root.Left, max.Val)
	}

	// get balanced factor of node
	root.Bf = GetBalancedFactor(root)

	// already balanced
	if root.Bf == 0 || root.Bf == -1 || root.Bf == 1 {
		return root
	}

	// left left imbalanced
	if root.Bf > 1 && root.Left.Bf > 0 {
		return RightRotation(root)
	}

	// right right imbalanced
	if root.Bf < -1 && root.Right.Bf < 0 {
		return LeftRotation(root)
	}

	// left right imbalanced
	if root.Bf > 1 && root.Left.Bf < 0 {
		root.Left = LeftRotation(root.Left)
		return RightRotation(root)
	}

	// right left imbalanced
	if root.Bf < -1 && root.Right.Bf > 0 {
		root.Right = RightRotation(root.Right)
		return LeftRotation(root)
	}

	return root
}

// LeftRotation rotate a node to left side
func LeftRotation(root *Node) *Node {
	x := root.Right
	y := x.Left

	// rotate the root to left side
	x.Left = root
	root.Right = y

	// calculate the balanced factor of root and x
	root.Bf = GetBalancedFactor(root)
	x.Bf = GetBalancedFactor(x)

	return x
}

// RightRotation rotate a node to right side
func RightRotation(root *Node) *Node {
	x := root.Left
	y := x.Right

	// rotate the root to right side
	x.Right = root
	root.Left = y

	// calculate the balanced factor of root and x
	root.Bf = GetBalancedFactor(root)
	x.Bf = GetBalancedFactor(x)

	return x
}

// GetBalancedFactor gets a balanced factor of a node
func GetBalancedFactor(root *Node) int {
	hleft := GetHeightRecursive(root.Left)
	hright := GetHeightRecursive(root.Right)
	root.Bf = hleft - hright

	return root.Bf
}

// GetHeightRecursive gets a height of a node
func GetHeightRecursive(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(
		float64(GetHeightRecursive(root.Left)),
		float64(GetHeightRecursive(root.Right))))
}

// LevelOrder Print tree via levels
func LevelOrder(root *Node) {
	height := GetHeightRecursive(root)

	for h := 1; h <= height; h++ {
		GivenLevelOrderRecursive(root, h)
		fmt.Println()
	}
}

// GivenLevelOrderRecursive print specific level
func GivenLevelOrderRecursive(root *Node, level int) {
	if level == 1 && root != nil {
		fmt.Printf("(%v,%v) ", root.Val, root.Bf)
		return
	}

	if level == 1 && root == nil {
		fmt.Print("(no,no)", " ")
		return
	}

	if root != nil && root.Left != nil {
		GivenLevelOrderRecursive(root.Left, level-1)
	} else {
		GivenLevelOrderRecursive(nil, level-1)
	}
	if root != nil && root.Right != nil {
		GivenLevelOrderRecursive(root.Right, level-1)
	} else {
		GivenLevelOrderRecursive(nil, level-1)
	}
}
