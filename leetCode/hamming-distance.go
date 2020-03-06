// https://leetcode.com/problems/hamming-distance
// Just sol to the problem, It does not include the I/O part

func hammingDistance(x int, y int) int {
    count := 0
    tmp := x ^ y

    for tmp > 0 {
        if tmp & 1 == 1 { count++ }
        tmp = tmp >> 1
    }
    return count
}
