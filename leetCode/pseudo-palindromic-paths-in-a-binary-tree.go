package main

func pseudoPalindromicPaths(root *TreeNode) int {
	var res int
	preorder(root, &res, [10]int{})
	return res
}

func preorder(root *TreeNode, res *int, parents [10]int) {
	if root == nil {
		return
	}

	parents[root.Val]++

	if root.Left == nil && root.Right == nil {
		*res++
		odd := false
		for i := 1; i < 10; i++ {
			if parents[i]&1 == 1 {
				if odd == true {
					*res--
					break
				}
				odd = true
			}
		}
	}

	preorder(root.Left, res, parents)
	preorder(root.Right, res, parents)
}
