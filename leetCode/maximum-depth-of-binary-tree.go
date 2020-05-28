package main

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lHeight := maxDepth(root.Left)
	rHeight := maxDepth(root.Right)

	return int(math.Max(float64(lHeight), float64(rHeight))) + 1
}
