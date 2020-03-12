// https://leetcode.com/problems/length-of-last-word
// Just sol to the problem, It does not include the I/O part

func lengthOfLastWord(s string) int {

    res := 0
    l := len(s)

    // remove length of spaces 
    for i:=l-1; i>=0; i-- {
        if s[i] == ' ' {
            l--
        } else {
            break
        }
    }

    // get the length of word
    for i:=l-1; i>=0; i-- {
        if s[i] == ' ' { break }
        res++
    }

    return res
}
