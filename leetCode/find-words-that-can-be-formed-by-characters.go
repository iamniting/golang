// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters
// Just sol to the problem, It does not include the I/O part

func countCharacters(words []string, chars string) int {

    m := make(map[rune]int)
    res := 0

    // get frequency of chars
    for _, c := range chars {
        m[c]++
    }

    for _, word := range words {
        // copy a map into tmp map
        tmpMap := make(map[rune]int)
        for k, v := range m {
            tmpMap[k] = v
        }

        // go through the word and check if present in tmp map
        isGood := true
        for _, c := range word {
            if val, ok := tmpMap[c]; !ok || val == 0 {
                isGood = false
                break
            }
            tmpMap[c]--
        }
        if isGood { res += len(word) }
    }

    return res
}
