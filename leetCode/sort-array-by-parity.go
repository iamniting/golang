// https://leetcode.com/problems/sort-array-by-parity
// Just sol to the problem, It does not include the I/O part

func sortArrayByParity(A []int) []int {

    even, odd := []int{}, []int{}

    for _, i := range A {
        if i & 1 == 1 {
            odd = append(odd, i)
        } else {
            even = append(even, i)
        }
    }

    return append(even, odd...)
}
