package main

func goodNodes(root *TreeNode) int {
	var res int
	preOrder(root, root, &res)
	return res
}

func preOrder(root, max *TreeNode, res *int) {

	if root == nil {
		return
	}

	if root.Val >= max.Val {
		*res++
		max = root
	}

	preOrder(root.Left, max, res)
	preOrder(root.Right, max, res)
}
