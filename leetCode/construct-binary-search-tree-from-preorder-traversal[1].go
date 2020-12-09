package main

import "math"

func bstFromPreorder(preorder []int) *TreeNode {

	min, max, index := math.MaxInt32, math.MinInt32, 0

	for _, val := range preorder {
		min = int(math.Min(float64(min), float64(val)))
		max = int(math.Max(float64(max), float64(val)))
	}

	return buildBST(preorder, &index, min, max)
}

func buildBST(preorder []int, index *int, min, max int) *TreeNode {

	// if all elements are done ||
	// val is less than min, which means can not add val to subtree ||
	// val is more than max, which means can not add val to subtree
	if *index >= len(preorder) || preorder[*index] < min || preorder[*index] > max {
		return nil
	}

	root := &TreeNode{preorder[*index], nil, nil}
	*index++

	root.Left = buildBST(preorder, index, min, root.Val)
	root.Right = buildBST(preorder, index, root.Val, max)

	return root
}
