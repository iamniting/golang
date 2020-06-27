package main

func isCousins(root *TreeNode, x int, y int) bool {
	var px, py, dx, dy int
	var preorder func(root *TreeNode, parent, dep int)

	preorder = func(root *TreeNode, parent, dep int) {
		if root == nil {
			return
		}
		if root.Val == x {
			dx = dep
			px = parent
		}
		if root.Val == y {
			dy = dep
			py = parent
		}
		if px == 0 || py == 0 {
			preorder(root.Left, root.Val, dep+1)
			preorder(root.Right, root.Val, dep+1)
		}
	}

	preorder(root, root.Val, 0)
	return dx == dy && px != py
}
