package main

import "math"

func minDiffInBST(root *TreeNode) int {
	res, prev := math.MaxInt32, math.MinInt32

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)

		if prev != math.MinInt32 {
			res = int(math.Min(float64(res), float64(root.Val-prev)))
		}

		prev = root.Val
		inorder(root.Right)
	}

	inorder(root)
	return res
}
