// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts
// Just sol to the problem, It does not include the I/O part

func generateTheString(n int) string {

    // odd num
    if n & 1 == 1 {
        return strings.Repeat("a", n)
    }

    // even num
    return strings.Repeat("a", n-1) + "b"
}
