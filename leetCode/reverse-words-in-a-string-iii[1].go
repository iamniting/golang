// https://leetcode.com/problems/reverse-words-in-a-string-iii
// Just sol to the problem, It does not include the I/O part

func reverseWords(s string) string {
    // helper func to reverse a word
    reverse := func(word string) string {
        ans := ""
        for i:=len(word)-1; i>=0; i-- {
            ans += string(word[i])
        }
        return ans
    }

    res, index := "", 0
    for i, c := range s {
        if c == ' ' {
            res += reverse(s[index:i]) + string(' ')
            index = i+1
        // if reaches at end
        } else if i == len(s) - 1 {
            res += reverse(s[index:])
        }
    }
    return res
}
