// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
// Just sol to the problem, It does not include the I/O part

func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, candie := range candies {
        if max < candie {
            max = candie
        }
    }

    res := make([]bool, len(candies))
    for i, candie := range candies {
        if candie + extraCandies >= max {
            res[i] = true
        }
    }
    return res
}
