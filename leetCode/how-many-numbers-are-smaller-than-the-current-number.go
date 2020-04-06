// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number
// Just sol to the problem, It does not include the I/O part

func smallerNumbersThanCurrent(nums []int) []int {

    res := make([]int, len(nums))

    for i, _ := range nums {

        for j, _ := range nums {

            if nums[i] > nums[j] {
                res[i]++
            }
        }
    }

    return res
}
