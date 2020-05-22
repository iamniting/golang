// https://leetcode.com/problems/univalued-binary-tree
// Just sol to the problem, It does not include the I/O part

func isUnivalTree(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Left.Val != root.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val != root.Val {
			return false
		}
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
