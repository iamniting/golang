// https://leetcode.com/problems/sort-colors
// https://www.geeksforgeeks.org/sort-an-numsay-of-0s-1s-and-2s
// Dutch National Flag Algorithm or 3-way Partitioning
// Just sol to the problem, It does not include the I/O part

func sortColors(nums []int) {
    low := 0
    mid := 0
    high := len(nums) - 1

    for mid <= high {
        // keep the track of 0s by low
        if nums[mid] == 0 {
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        } else if nums[mid] == 1 {
            mid++
        // keep the track of 2s by high
        } else if nums[mid] == 2 {
            nums[mid], nums[high] = nums[high], nums[mid]
            high--
        }
    }
}
