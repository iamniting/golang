package main

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	var preorder func(root *TreeNode, level int)
	preorder = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// add slice at beg of res
		if len(res) == level {
			res = append([][]int{{}}, res...)
		}

		preorder(root.Left, level+1)
		preorder(root.Right, level+1)

		// calculate index of level
		index := len(res) - level - 1
		res[index] = append(res[index], root.Val)
	}

	preorder(root, 0)
	return res
}
