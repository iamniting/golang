// https://leetcode.com/problems/longest-palindromic-substring
// Just sol to the problem, It does not include the I/O part

func longestPalindrome(s string) string {
    palin := ""
    res := ""

    for i, _ := range s {

        // odd length palindrome
        l := i-1
        r := i+1

        for l > -1 && r < len(s) && s[l] == s[r] {
            l--
            r++
        }

        palin = s[l+1:r]

        if len(palin) > len(res) { res = palin }

        // even length palindrome
        l = i
        r = i+1

        for l > -1 && r < len(s) && s[l] == s[r] {
            l--
            r++
        }
        palin = s[l+1:r]

        if len(palin) > len(res) { res = palin }
    }
    return res
}
