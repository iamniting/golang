package main

import (
	"fmt"
	"math"
)

func findSecondMinimumValue(root *TreeNode) int {

	a, b := math.MaxInt32+1, math.MaxInt32+1

	preorder(root, &a, &b)

	if a == math.MaxInt32+1 || b == math.MaxInt32+1 {
		return -1
	}
	return b
}

func preorder(root *TreeNode, a, b *int) {

	if root == nil {
		return
	}

	if root.Val < *a {
		*a, *b = root.Val, *a
	} else if root.Val < *b && root.Val > *a {
		*b = root.Val
	}

	preorder(root.Left, a, b)
	preorder(root.Right, a, b)
}
