// https://leetcode.com/problems/remove-element
// Just sol to the problem, It does not include the I/O part

func removeElement(nums []int, val int) int {

    i, j := 0, len(nums) - 1

    for i <= j {

        if nums[i] == val {
            nums[i] , nums[j] = nums[j], nums[i]
            j--
            continue
        }

        i++
    }
    return i
}
