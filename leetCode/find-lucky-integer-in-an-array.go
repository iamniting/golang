// https://leetcode.com/problems/find-lucky-integer-in-an-array
// Just sol to the problem, It does not include the I/O part

func findLucky(arr []int) int {

    m := make(map[int]int)

    for _, n := range arr {
        m[n]++
    }

    res := -1

    for k, v := range m {
        if k == v && k > res{
            res = k
        }
    }
    return res
}
