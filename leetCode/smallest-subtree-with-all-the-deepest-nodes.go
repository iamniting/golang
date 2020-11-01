package main

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, res := postOrder(root)
	return res
}

func postOrder(root *TreeNode) (int, *TreeNode) {

	if root == nil {
		return 0, nil
	}

	lh, left := postOrder(root.Left)
	rh, right := postOrder(root.Right)

	if lh > rh {
		return lh + 1, left
	}

	if rh > lh {
		return rh + 1, right
	}

	// if lh == rh means common ancestor
	return rh + 1, root
}
