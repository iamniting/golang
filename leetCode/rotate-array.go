// https://leetcode.com/problems/rotate-array
// Just sol to the problem, It does not include the I/O part

func rotate(nums []int, k int)  {

    l := len(nums)
    k = k % l

    // get rotated array in tmp var
    tmp := append(nums[l-k:], nums[:l-k]...)
    copy(nums, tmp)
}
