// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix
// Just sol to the problem, It does not include the I/O part

func countNegatives(grid [][]int) int {

    res := 0

    for _, m := range grid {
        for _, n := range m {

            if n < 0 { res++ }
        }
    }

    return res
}
