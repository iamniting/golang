package main

import "math"

func findTilt(root *TreeNode) int {
	var res int
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	l := postorder(root.Left, res)
	r := postorder(root.Right, res)

	*res += int(math.Abs(float64(l - r)))

	return l + r + root.Val
}
