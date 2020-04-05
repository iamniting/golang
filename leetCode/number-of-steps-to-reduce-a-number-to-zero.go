// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero
// Just sol to the problem, It does not include the I/O part

func numberOfSteps (num int) int {

    res := 0

    for num > 0 {
        if num & 1 == 1 {
            num--
        } else {
            num >>= 1
        }
        res++
    }

    return res
}
