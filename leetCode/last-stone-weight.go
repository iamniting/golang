// https://leetcode.com/problems/last-stone-weight
// Just sol to the problem, It does not include the I/O part

func lastStoneWeight(stones []int) int {

    for len(stones) > 1 {
        // sort reverse
        sort.Sort(sort.Reverse(sort.IntSlice(stones)))
        x, y := stones[0], stones[1]
        // delete first two elements and add difference
        stones = append(stones[2:], x-y)
    }

    return stones[0]
}
