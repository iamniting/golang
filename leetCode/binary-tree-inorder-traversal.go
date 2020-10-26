package main

func inorderTraversal(root *TreeNode) []int {
	var slice []int
	inorder(root, &slice)
	return slice
}

func inorder(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, slice)

	*slice = append(*slice, root.Val)

	inorder(root.Right, slice)
}
