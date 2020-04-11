// https://leetcode.com/problems/maximum-69-number
// Just sol to the problem, It does not include the I/O part

func maximum69Number (num int) int {

    numStr := strconv.Itoa(num)

    for i, c := range numStr {
        if c == '6' {
            num, _ = strconv.Atoi(numStr[:i] + "9" + numStr[i+1:])
            return num
        }
    }
    return num
}
