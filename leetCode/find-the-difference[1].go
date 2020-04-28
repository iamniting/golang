// https://leetcode.com/problems/find-the-difference
// Just sol to the problem, It does not include the I/O part

func findTheDifference(s string, t string) byte {

    var res rune

    for _, c := range s + t {
        res ^= c
    }

    return byte(res)
}
