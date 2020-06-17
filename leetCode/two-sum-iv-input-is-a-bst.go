package main

func findTarget(root *TreeNode, k int) bool {
	var preorder func(root *TreeNode, m map[int]interface{}, k int) bool

	preorder = func(root *TreeNode, m map[int]interface{}, k int) bool {

		if root == nil {
			return false
		}

		if _, ok := m[k-root.Val]; ok {
			return true
		}

		m[root.Val] = nil
		return preorder(root.Left, m, k) || preorder(root.Right, m, k)
	}

	m := map[int]interface{}{}
	return preorder(root, m, k)
}
