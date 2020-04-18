// https://leetcode.com/problems/reverse-words-in-a-string-iii
// Just sol to the problem, It does not include the I/O part

func reverseWords(s string) string {

    // seprate words with spaces
    words := strings.Fields(s)
    res := ""

    // reverse words one by one and add it into res
    for _, word := range words {
        for i:=len(word)-1; i>=0; i-- {
            res += string(word[i])
        }
        res += " "
    }

    // remove last space
    return strings.TrimSpace(res)
}
