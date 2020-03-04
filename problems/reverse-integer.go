// https://leetcode.com/problems/reverse-integer
// Just sol to the problem, It does not include the I/O part

func reverse(x int) int {
    res := 0
    minus := false

    if x < 0 { minus = true }

    x = int(math.Abs(float64(x)))

    for x > 0 {
        mod := x % 10
        x = x / 10
        res = (res * 10) + mod
    }

    // overflow condition
    if res > 2147483648 { return 0 }

    if minus { return -res }

    return res
}
