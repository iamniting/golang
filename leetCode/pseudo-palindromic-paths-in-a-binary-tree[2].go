package main

func pseudoPalindromicPaths(root *TreeNode) int {
	return preorder(root, 0)
}

func preorder(root *TreeNode, path int) int {
	if root == nil {
		return 0
	}

	var res int
	// compute occurrences of each digit in path
	// The idea is to keep the frequency of digit 1 in the first bit, 2 in the second bit, etc
	path = path ^ (1 << root.Val)

	if root.Left == nil && root.Right == nil {
		// Now, to ensure that at most one digit has an odd frequency,
		// one has to check that path is a power of two, i.e., at most one bit is set to one.
		if path&(path-1) == 0 {
			res++
		}
	}

	return res + preorder(root.Left, path) + preorder(root.Right, path)
}
