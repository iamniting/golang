package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	preorder(root, "", &ans)
	return ans
}

func preorder(root *TreeNode, prefix string, ans *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		prefix += "->" + strconv.Itoa(root.Val)
		*ans = append(*ans, prefix[2:])
		return
	}

	preorder(root.Left, prefix+"->"+strconv.Itoa(root.Val), ans)
	preorder(root.Right, prefix+"->"+strconv.Itoa(root.Val), ans)
}
