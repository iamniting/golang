// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero
// Just sol to the problem, It does not include the I/O part

func numberOfSteps (num int) int {

    res := 0

    for num > 0 {
        // (num & 1) = 1 for the odd num which means if odd num +2 in res
        res += (num & 1) + 1
        num >>= 1
    }

    return res - 1
}
