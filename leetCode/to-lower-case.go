// https://leetcode.com/problems/to-lower-case
// Just sol to the problem, It does not include the I/O part

func toLowerCase(str string) string {

    res := ""

    for _, c := range str {

        res += string(c | ' ')
    }

    return res
}
