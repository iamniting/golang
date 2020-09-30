package main

import "math"

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 || r == 0 {
		return 1 + l + r
	}

	return 1 + int(math.Min(float64(l), float64(r)))
}
