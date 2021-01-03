package main

func pseudoPalindromicPaths(root *TreeNode) int {
	return preorder(root, [10]int{})
}

func preorder(root *TreeNode, parents [10]int) int {
	if root == nil {
		return 0
	}

	var res int
	parents[root.Val]++

	if root.Left == nil && root.Right == nil {
		res++
		odd := false
		for i := 1; i < 10; i++ {
			if parents[i]&1 == 1 {
				if odd == true {
					res--
					break
				}
				odd = true
			}
		}

	}

	res += preorder(root.Left, parents)
	res += preorder(root.Right, parents)
	return res
}
