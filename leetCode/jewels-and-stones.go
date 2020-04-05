// https://leetcode.com/problems/jewels-and-stones
// Just sol to the problem, It does not include the I/O part

func numJewelsInStones(J string, S string) int {

    res := 0

    for _, j := range J {
        for _, s := range S {
            if j == s {
                res++
            }
        }
    }

    return res
}
