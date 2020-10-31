package main

func smallestFromLeaf(root *TreeNode) string {

	// this sol is not working for one TC and IMO TC is wrong as TC does not satisfy for its subtree
	// TC [25,1,null,0,0,1,null,null,null,0] and expected ans is "ababz" my ans is "abz"
	if root == nil {
		return ""
	}

	rootVal := string('a' + root.Val)

	left := smallestFromLeaf(root.Left) + rootVal
	right := smallestFromLeaf(root.Right) + rootVal

	if left == rootVal {
		return right
	}

	if right == rootVal {
		return left
	}

	if left < right {
		return left
	}

	return right
}
