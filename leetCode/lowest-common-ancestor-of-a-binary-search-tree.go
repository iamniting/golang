package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// if p and q both are lower than root
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// if p and q both are greater than root
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	// else this is a split point of both
	return root
}
