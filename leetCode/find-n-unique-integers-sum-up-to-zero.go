// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero
// Just sol to the problem, It does not include the I/O part

func sumZero(n int) []int {

    res := make([]int, n)
    x := 1

    for i:=0; i<n-1; i=i+2 {
        res[i], res[i+1] = x, -x
        x++
    }

    return res
}
