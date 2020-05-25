package main

func increasingBST(root *TreeNode) *TreeNode {
	var inorder func(root *TreeNode, slice *[]*TreeNode)

	inorder = func(root *TreeNode, slice *[]*TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left, slice)
		*slice = append(*slice, root)
		inorder(root.Right, slice)
	}

	res := []*TreeNode{}
	inorder(root, &res)

	for i := 0; i < len(res)-1; i++ {
		res[i].Left = nil
		res[i].Right = res[i+1]
	}
	return res[0]
}
