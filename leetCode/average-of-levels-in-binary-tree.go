package main

func averageOfLevels(root *TreeNode) []float64 {
	var inorder func(root *TreeNode, level int, sum, count *[]int)
	inorder = func(root *TreeNode, level int, sum, count *[]int) {

		if root == nil {
			return
		}

		if len(*sum) > level {
			(*sum)[level] += root.Val
			(*count)[level]++
		} else {
			*sum = append(*sum, root.Val)
			*count = append(*count, 1)
		}

		inorder(root.Left, level+1, sum, count)
		inorder(root.Right, level+1, sum, count)
	}

	sum, count := []int{}, []int{}

	inorder(root, 0, &sum, &count)

	res := make([]float64, len(sum))
	for i, s := range sum {
		res[i] = float64(s) / float64(count[i])
	}

	return res
}
