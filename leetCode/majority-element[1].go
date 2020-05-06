// https://leetcode.com/problems/majority-element
// Just sol to the problem, It does not include the I/O part

func majorityElement(nums []int) int {
    res, count := 0, 0

    for _, n := range nums {
        if count == 0 {
            res = n
        }
        if res == n {
            count++
        } else {
            count--
        }
    }
    return res
}
