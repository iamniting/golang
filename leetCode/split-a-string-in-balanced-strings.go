// https://leetcode.com/problems/split-a-string-in-balanced-strings
// Just sol to the problem, It does not include the I/O part

func balancedStringSplit(s string) int {
    res, bal := 0, 0

    for _, c := range s {
        if c == 'L' {
            bal++
        } else {
            bal--
        }

        if bal == 0 {
            res++
        }
    }
    return res
}
