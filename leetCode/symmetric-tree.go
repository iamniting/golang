// https://leetcode.com/problems/symmetric-tree
// Just sol to the problem, It does not include the I/O part

func helper(l, r *TreeNode) bool {
    // if both roots are nil
    if l == nil && r == nil { return true }

    // if only one root is nil
    if l == nil || r == nil { return false }

    // if values are same
    return l.Val == r.Val && helper(l.Left, r.Right) && helper(l.Right, r.Left)
}

func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }

    return helper(root.Left, root.Right)
}
