package main

import "math"

func findMode(root *TreeNode) []int {

	res := []int{}
	var max, count, prev int
	prev = math.MinInt32

	preorder(root, &res, &prev, &count, &max)

	return res
}

func preorder(root *TreeNode, res *[]int, prev, count, max *int) {
	if root == nil {
		return
	}

	preorder(root.Left, res, prev, count, max)

	if root.Val == *prev {
		*count++
	} else {
		*count = 1
	}

	*prev = root.Val

	if *count > *max {
		*res = []int{root.Val}
		*max = *count
	} else if *count == *max {
		*res = append(*res, root.Val)
	}

	preorder(root.Right, res, prev, count, max)
}
