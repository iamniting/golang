// https://leetcode.com/problems/search-insert-position
// Just sol to the problem, It does not include the I/O part 

func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    for left <= right {
        mid := (left + right) / 2

        if nums[mid] == target { return mid }

        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}
