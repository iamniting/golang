// https://leetcode.com/problems/string-matching-in-an-array
// Just sol to the problem, It does not include the I/O part

func stringMatching(words []string) []string {

    resMap := make(map[string]bool)
    res := []string{}

    for i, word := range words {
        for j, wrd := range words {
            if i != j && strings.Contains(wrd, word) {
                resMap[word] = true
            }
        }
    }

    for key := range resMap {
        res = append(res, key)
    }

    return res
}
