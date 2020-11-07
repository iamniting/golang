package main

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	var maxNum, maxIndex int

	for i, n := range nums {
		if n > maxNum {
			maxIndex, maxNum = i, n
		}
	}

	root := &TreeNode{maxNum, nil, nil}

	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}
