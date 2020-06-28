package main

func levelOrderBottom(root *TreeNode) [][]int {

	var preorder func(root *TreeNode, level int, m map[int][]int)
	preorder = func(root *TreeNode, level int, m map[int][]int) {
		if root == nil {
			return
		}

		m[level] = append(m[level], root.Val)

		preorder(root.Left, level+1, m)
		preorder(root.Right, level+1, m)
	}

	m := map[int][]int{}
	res := [][]int{}

	preorder(root, 1, m)

	for i := len(m); i > 0; i-- {
		res = append(res, m[i])
	}
	return res
}
