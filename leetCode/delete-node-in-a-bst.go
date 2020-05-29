package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		// if leaf node which need to be deleted
	} else if root.Val == key && root.Left == nil && root.Right == nil {
		return nil
		// if node to be deleted having one or both child
	} else {
		// node with only left child
		if root.Left == nil {
			return root.Right
			// node with only right child
		} else if root.Right == nil {
			return root.Left
		}

		// get max value of left sub tree
		max := root.Left
		for max.Right != nil {
			max = max.Right
		}

		// replace root.Val with the max value
		root.Val = max.Val

		// delete max value of left sub tree
		root.Left = deleteNode(root.Left, max.Val)
	}
	return root
}
