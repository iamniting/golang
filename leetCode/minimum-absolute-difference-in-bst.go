package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	var inorder func(root *TreeNode, p *[]int)
	inorder = func(root *TreeNode, p *[]int) {
		if root == nil {
			return
		}
		inorder(root.Left, p)
		*p = append(*p, root.Val)
		inorder(root.Right, p)
	}

	slice := []int{}
	res := math.MaxInt32

	inorder(root, &slice)

	for i := 1; i < len(slice); i++ {
		res = int(math.Min(float64(res), float64(slice[i]-slice[i-1])))
	}

	return res
}
