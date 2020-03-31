// https://leetcode.com/problems/number-of-1-bits
// Just sol to the problem, It does not include the I/O part

func hammingWeight(num uint32) int {
    res := 0

    for num > 0 {
        if num & 1 == 1 {
            res++
        }
        num = num >> 1
    }
    return res
}
