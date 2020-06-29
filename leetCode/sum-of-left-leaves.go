package main

func sumOfLeftLeaves(root *TreeNode) int {

	var preorder func(root *TreeNode, isLeft bool, res *int)
	preorder = func(root *TreeNode, isLeft bool, res *int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil && isLeft {
			*res += root.Val
		}

		preorder(root.Left, true, res)
		preorder(root.Right, false, res)
	}

	var res int
	preorder(root, false, &res)

	return res
}
