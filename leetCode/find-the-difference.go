// https://leetcode.com/problems/find-the-difference
// Just sol to the problem, It does not include the I/O part

func findTheDifference(s string, t string) byte {

    m1 := make(map[rune]int)
    m2 := make(map[rune]int)

    for _, c := range s {
        m1[c]++
    }

    for _, c := range t {
        m2[c]++
    }

    for _ ,c := range t {
        if m1[c] != m2[c] {
            return byte(c)
        }
    }
    return byte(0)
}
