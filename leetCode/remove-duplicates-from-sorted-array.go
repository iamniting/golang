// https://leetcode.com/problems/remove-duplicates-from-sorted-array
// Just sol to the problem, It does not include the I/O part

func removeDuplicates(nums []int) int {
    index := 0

    for i := 1; i < len(nums); i++ {

        if nums[index] != nums[i] {
            index++
            nums[index] = nums[i]
        }
    }

    return index+1
}
