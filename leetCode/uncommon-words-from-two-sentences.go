// https://leetcode.com/problems/uncommon-words-from-two-sentences
// Just sol to the problem, It does not include the I/O part

func uncommonFromSentences(A string, B string) []string {

    res := []string{}
    m := make(map[string]int)

    // store words and occurance
    for _, word := range strings.Fields(A + " " + B) {
        m[word]++
    }

    // if occurance is 1 append it into res
    for k, v := range m {
        if v == 1 {
            res = append(res, k)
        }
    }
    return res
}
