package main

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	ans := 0
	if root.Val >= L && root.Val <= R {
		ans += root.Val
	}

	if root.Val > L {
		ans += rangeSumBST(root.Left, L, R)
	}
	if root.Val < R {
		ans += rangeSumBST(root.Right, L, R)
	}
	return ans
}
