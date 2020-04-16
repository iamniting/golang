// https://leetcode.com/problems/height-checker
// Just sol to the problem, It does not include the I/O part

func heightChecker(heights []int) int {

    sorted := make([]int, len(heights))
    copy(sorted[:], heights)
    sort.Ints(sorted)

    res := 0
    for i, _ := range heights {
        if sorted[i] != heights[i] {
            res++
        }
    }

    return res
}
