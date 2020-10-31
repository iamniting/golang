package main

func smallestFromLeaf(root *TreeNode) string {

	return postOrder(root, "")
}

func postOrder(root *TreeNode, parent string) string {
	if root == nil {
		return ""
	}

	rootVal := string('a' + root.Val)

	left := postOrder(root.Left, rootVal) + rootVal
	right := postOrder(root.Right, rootVal) + rootVal

	if left == rootVal {
		return right
	}

	if right == rootVal {
		return left
	}

	if left+parent < right+parent {
		return left
	}

	return right
}
