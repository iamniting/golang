// https://leetcode.com/problems/two-sum/
// Just sol to the problem, It does not include the I/O part

func twoSum(nums []int, target int) []int {

    m := make(map[int]int)

    for i, n := range nums {
        tmp := target - n

        if index, ok := m[tmp]; ok {
            return []int{index, i}
        }
        m[n] = i
    }
    return []int{}
}
