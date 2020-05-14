// https://leetcode.com/problems/n-repeated-element-in-size-2n-array
// Just sol to the problem, It does not include the I/O part

func repeatedNTimes(A []int) int {
    m := make(map[int]bool)

    for _, n := range A {
        if m[n] {
            return n
        }
        m[n] = true
    }
    return 0
}
