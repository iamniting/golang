// https://leetcode.com/problems/single-number
// Just sol to the problem, It does not include the I/O part

func singleNumber(nums []int) int {

    res := 0

    // xor gives 0 when doing with same num
    // xor (2 ^ 2) = 0
    // xor (2 ^ 2 ^ 3) = 3
    // this way we will get the single num
    for _, n := range nums {
        res ^= n
    }

    return res
}
