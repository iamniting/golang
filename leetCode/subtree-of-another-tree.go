package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val && isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
