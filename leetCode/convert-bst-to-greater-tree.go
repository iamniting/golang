package main

func convertBST(root *TreeNode) *TreeNode {

	var inorder func(root *TreeNode, sum *int)
	inorder = func(root *TreeNode, sum *int) {
		if root == nil {
			return
		}

		inorder(root.Right, sum)

		root.Val += *sum
		*sum = root.Val

		inorder(root.Left, sum)
	}

	sum := 0
	inorder(root, &sum)
	return root
}
