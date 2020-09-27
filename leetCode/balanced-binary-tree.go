package main

import "math"

func isBalanced(root *TreeNode) bool {
	res := true
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *bool) float64 {
	if root == nil {
		return 0
	}

	lh := postOrder(root.Left, res)
	rh := postOrder(root.Right, res)

	if math.Abs(lh-rh) > 1 {
		*res = false
	}

	return math.Max(lh, rh) + 1
}
