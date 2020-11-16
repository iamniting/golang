package main

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	var slice1, slice2, res []int

	inorder(root1, &slice1)
	inorder(root2, &slice2)

	var i, j int

	for i < len(slice1) || j < len(slice2) {
		if i < len(slice1) && j < len(slice2) {
			if slice1[i] < slice2[j] {
				res = append(res, slice1[i])
				i++
			} else {
				res = append(res, slice2[j])
				j++
			}
		} else if i < len(slice1) {
			res = append(res, slice1[i:]...)
			i = len(slice1)
		} else {
			res = append(res, slice2[j:]...)
			j = len(slice2)
		}
	}

	return res
}

func inorder(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, slice)
	*slice = append(*slice, root.Val)
	inorder(root.Right, slice)
}
