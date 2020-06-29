package main

func sumOfLeftLeaves(root *TreeNode) int {

	var sum int

	if root == nil {
		return sum
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
