// https://leetcode.com/problems/unique-number-of-occurrences
// Just sol to the problem, It does not include the I/O part

func uniqueOccurrences(arr []int) bool {

    // get the occurance count
    m := make(map[int]int)

    for _, n := range arr {
        m[n]++
    }

    // get the values in the set
    set := make(map[int]bool)

    for _, v := range m {
        set[v] = true
    }

    return len(set) == len(m)
}
