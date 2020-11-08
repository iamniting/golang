package main

import "math"

func longestUnivaluePath(root *TreeNode) int {
	var res int
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := postOrder(root.Left, res)
	right := postOrder(root.Right, res)

	var lsum, rsum int

	if root.Left != nil && root.Left.Val == root.Val {
		lsum = left + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		rsum = right + 1
	}

	*res = int(math.Max(float64(*res), float64(lsum+rsum)))
	return int(math.Max(float64(lsum), float64(rsum)))
}
