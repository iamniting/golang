// https://leetcode.com/problems/maximum-69-number
// Just sol to the problem, It does not include the I/O part

func maximum69Number (num int) int {
    res := 0
    divsr := 1
    flag := true

    // calculate accurate divisor for num
    for divsr < num { divsr *= 10 }
    divsr /= 10

    // calculate res digit by digit and change first ever 6 digit
    for num > 0 {

        quot := num / divsr
        num %= divsr
        divsr /= 10

        if flag && quot == 6 {
            flag = false
            quot = 9
        }

        res = (res * 10) + quot
    }
    return res
}
