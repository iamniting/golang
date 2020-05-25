package main

func sumRootToLeaf(root *TreeNode) int {

	var preOrder func(root *TreeNode, n int) int

	preOrder = func(root *TreeNode, n int) int {
		if root == nil {
			return 0
		}

		n = (n << 1) | root.Val

		if root.Left == root.Right {
			return n
		}

		return preOrder(root.Left, n) + preOrder(root.Right, n)
	}

	return preOrder(root, 0)
}
