// https://leetcode.com/problems/jewels-and-stones
// Just sol to the problem, It does not include the I/O part

func numJewelsInStones(J string, S string) int {

    res := 0
    type void struct {}
    m := make(map[rune]void)

    for _, j := range J {

        m[j] = void{}
    }

    for _, s := range S {

        if _, ok := m[s]; ok {
            res++
        }
    }

    return res
}
