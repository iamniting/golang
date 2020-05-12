// https://leetcode.com/problems/increasing-decreasing-string
// Just sol to the problem, It does not include the I/O part

func sortString(s string) string {
    freq := make([]byte, 26)
    res := make([]byte, len(s))

    for _, c := range s {
        freq[c-'a']++
    }

    index := 0
    for index < len(s) {
        for i, n := range freq {
            if n > 0 {
                res[index] = byte(i + 'a')
                freq[i]--
                index++
            }
        }
        for i:=25; i>-1; i-- {
            if freq[i] > 0 {
                res[index] = byte(i + 'a')
                freq[i]--
                index++
            }
        }
    }
    return string(res)
}
