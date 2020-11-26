package main

func goodNodes(root *TreeNode) int {
	return preOrder(root, root)
}

func preOrder(root, max *TreeNode) int {

	if root == nil {
		return 0
	}

	var count int
	if root.Val >= max.Val {
		count++
		max = root
	}

	return count + preOrder(root.Left, max) + preOrder(root.Right, max)
}
