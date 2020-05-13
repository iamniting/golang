// https://leetcode.com/problems/self-dividing-numbers
// Just sol to the problem, It does not include the I/O part

func selfDividingNumbers(left int, right int) []int {

    IsSelfDiv := func(n int) bool {
        s := strconv.Itoa(n)
        for _, c := range s {
            if c == '0' { return false }
            if n % int(c - '0') != 0 { return false }
        }
        return true
    }

    res := []int{}
    for i:=left; i<=right; i++ {
        if IsSelfDiv(i) {
            res = append(res, i)
        }
    }
    return res
}
