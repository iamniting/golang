// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum
// Just sol to the problem, It does not include the I/O part

func minStartValue(nums []int) int {
    ans, sum := 0, 0

    // get minimum required num to make sum always greater than 0
    for _, n := range nums {
        sum += n
        ans = int(math.Min(float64(ans), float64(sum)))
    }
    // change the ans to positive and add +1 so that sum will not be zero
    return int(math.Abs(float64(ans))) + 1
}
