// https://leetcode.com/problems/implement-strstr
// Just sol to the problem, It does not include the I/O part

func strStr(haystack string, needle string) int {

    if needle == "" { return 0 }

    for i := 0; i <= len(haystack) - len(needle); i++ {

        j := 0
        for ; j < len(needle); j++ {
            if haystack[i+j] != needle[j] {
                break
            }
        }

        if j == len(needle) { return i }
    }

    return -1
}
