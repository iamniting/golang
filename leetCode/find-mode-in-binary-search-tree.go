package main

func findMode(root *TreeNode) []int {

	m := map[int]int{}
	res := []int{}
	var min int

	inorder(root, m)

	for val, occ := range m {
		if min < occ {
			res = []int{val}
			min = occ
		} else if min == occ {
			res = append(res, val)
		}
	}

	return res
}

func inorder(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	inorder(root.Left, m)
	m[root.Val]++
	inorder(root.Right, m)
}
