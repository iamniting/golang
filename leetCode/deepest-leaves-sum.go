package main

func deepestLeavesSum(root *TreeNode) int {
	var res, level int
	preOrder(root, 1, &level, &res)
	return res
}

func preOrder(root *TreeNode, height int, level, res *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if height > *level {
			*level = height
			*res = root.Val
		} else if height == *level {
			*res += root.Val
		}
		return
	}

	preOrder(root.Left, height+1, level, res)
	preOrder(root.Right, height+1, level, res)
}
