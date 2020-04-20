// https://leetcode.com/problems/minimum-absolute-difference
// Just sol to the problem, It does not include the I/O part

func minimumAbsDifference(arr []int) [][]int {

    min := math.MaxInt32
    res := [][]int{}

    sort.Ints(arr)

    for i:=1; i <len(arr); i++ {
        diff := arr[i] - arr[i-1]
        // get minimum
        if diff < min {
            min = diff
            // if got new min make the res empty
            res = [][]int{}
        }

        // if min == diff update the res
        if min == diff {
            res = append(res, []int{arr[i-1], arr[i]})
        }
    }

    return res
}
