package main

func isSameTree(p *TreeNode, q *TreeNode) bool {

	// if both are nil
	if p == nil && q == nil {
		return true
	}

	// if one is nil and one have data
	if p == nil || q == nil {
		return false
	}

	// if roots values are same
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
