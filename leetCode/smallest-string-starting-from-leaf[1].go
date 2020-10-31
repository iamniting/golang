package main

func smallestFromLeaf(root *TreeNode) string {

	slice := postOrder(root)

	var index int
	for i, str := range slice {
		if str < slice[index] {
			index = i
		}
	}

	return slice[index]
}

func postOrder(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var slice []string
	rootVal := string('a' + root.Val)

	left := postOrder(root.Left)
	right := postOrder(root.Right)

	if root.Left == nil && root.Right == nil {
		slice = append(slice, rootVal)
		return []string{rootVal}
	}

	for _, s := range left {
		slice = append(slice, s+rootVal)
	}
	for _, s := range right {
		slice = append(slice, s+rootVal)
	}

	return slice
}
