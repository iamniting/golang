package main

func rangeSumBST(root *TreeNode, L int, R int) int {
	ans := 0
	preOrderRecursive(root, L, R, &ans)
	return ans
}

func preOrderRecursive(root *TreeNode, L, R int, ans *int) {
	if root == nil {
		return
	}
	if root.Val >= L && root.Val <= R {
		*ans += root.Val
	}
	preOrderRecursive(root.Left, L, R, ans)
	preOrderRecursive(root.Right, L, R, ans)
}
