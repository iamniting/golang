// https://leetcode.com/problems/reverse-string
// Just sol to the problem, It does not include the I/O part

func reverseString(s []byte) {

    i, j := 0, len(s)-1
    for i < j {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}
