package main

import "math"

func diameterOfBinaryTree(root *TreeNode) int {

	var res int

	postOrder(root, &res)

	return res
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	lh := postOrder(root.Left, res)
	rh := postOrder(root.Right, res)

	*res = int(math.Max(float64(*res), float64(lh+rh)))
	return 1 + int(math.Max(float64(lh), float64(rh)))
}
