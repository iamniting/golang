// https://leetcode.com/problems/majority-element
// Just sol to the problem, It does not include the I/O part

func majorityElement(nums []int) int {

    res := nums[0]
    m := make(map[int]int)

    // store count of elements in map
    for _, i := range nums {

        if _, ok := m[i]; ok {
            m[i]++
        } else {
            m[i] = 1
        }
    }

    // get the key with highest value
    for key := range m {

        if m[key] > m[res] {
            res = key
        }
    }

    return res
}
