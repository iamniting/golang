// https://leetcode.com/problems/find-numbers-with-even-number-of-digits
// Just sol to the problem, It does not include the I/O part

func findNumbers(nums []int) int {

    res := 0

    for _, n := range nums {

        count := 0

        for ;n > 0; n /= 10 { count++ }

        if count & 1 == 0 { res++ }
    }

    return res
}
