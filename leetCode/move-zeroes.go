// https://leetcode.com/problems/move-zeroes
// Just sol to the problem, It does not include the I/O part

func moveZeroes(nums []int) {
    n := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[n], nums[i] = nums[i], nums[n]
            n++
        }
    }
}
