package main

func bstFromPreorder(preorder []int) *TreeNode {

	var root *TreeNode

	for _, val := range preorder {
		root = insertIntoBST(root, val)
	}

	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{val, nil, nil}
	} else if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
