package main

func sumEvenGrandparent(root *TreeNode) int {
	return postOrder(root, nil, nil)
}

func postOrder(root, parent, gparent *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int

	sum += postOrder(root.Left, root, parent)
	sum += postOrder(root.Right, root, parent)

	if gparent != nil && gparent.Val&1 == 0 {
		sum += root.Val
	}

	return sum
}
