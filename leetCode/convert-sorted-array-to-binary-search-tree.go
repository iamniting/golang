// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree
// Just sol to the problem, It does not include the I/O part

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }

    mid := len(nums) / 2

    return &TreeNode{
        Val: nums[mid],
        Left: sortedArrayToBST(nums[:mid]),
        Right: sortedArrayToBST(nums[mid+1:]),
    }
}
