// https://leetcode.com/problems/shortest-distance-to-a-character
// Just sol to the problem, It does not include the I/O part

func shortestToChar(S string, C byte) []int {

    res := make([]int, len(S))
    // check forward
    index := math.MinInt32
    for i := range S {
        if S[i] == C {
            index = i
        }
        res[i] = i - index
    }
    // check backward
    index = math.MaxInt32
    for i:=len(S)-1; i>=0; i-- {
        if S[i] == C {
            index = i
        }
        // if res is less than prev res than change it
        if res[i] > index - i {
            res[i] = index - i
        }
    }
    return res
}
