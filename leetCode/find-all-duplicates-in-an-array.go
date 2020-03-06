// https://leetcode.com/problems/find-all-duplicates-in-an-array
// Just sol to the problem, It does not include the I/O part

func findDuplicates(nums []int) []int {
    res := []int{}

    for _, n := range nums {
        num := int(math.Abs(float64(n)))

        // if num is already occured add it in res slice
        if nums[num-1] < 0 {
            res = append(res, num)
        // if num is coming for first time mark it as visited with -num
        } else {
            nums[num-1] = -nums[num-1]
        }
    }

    return res
}
