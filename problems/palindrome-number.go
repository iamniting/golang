// https://leetcode.com/problems/palindrome-number
// Just sol to the problem, It does not include the I/O part

func isPalindrome(x int) bool {
    tmp := x
    res := 0

    for tmp > 0 {
        mod := tmp % 10
        tmp = tmp / 10
        res = (res * 10) + mod
    }

    if x == res {return true}
    return false
}
