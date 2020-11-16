package main

import "sort"

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	var res []int

	inorder(root1, &res)
	inorder(root2, &res)

	sort.Ints(res)

	return res
}

func inorder(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, slice)
	*slice = append(*slice, root.Val)
	inorder(root.Right, slice)
}
