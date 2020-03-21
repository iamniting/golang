// https://leetcode.com/problems/valid-palindrome
// Just sol to the problem, It does not include the I/O part

func isPalindrome(s string) bool {
    i, j := 0, len(s) - 1

    for i < j {

        for i < j && !isValid(s[i]) {
            i++
        }

        for i < j && !isValid(s[j]) {
            j--
        }

        if i < j && (s[i] | ' ') != (s[j] | ' ') {
            return false
        }
        i++; j--
    }
    return true
}

func isValid(s byte) bool {
    return (s >= '0' && s <= '9')
        || (s >= 'A' && s <= 'Z')
        || (s >= 'a' && s <= 'z')
}
