// https://leetcode.com/problems/integer-to-roman
// Just sol to the problem, It does not include the I/O part

func intToRoman(num int) string {
    values := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4 , 1}
    romans := []string{
        "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    for i, v := range values {

        if num >= v {
            // get repetition of romans
            r := num / v
            num = num % v
            res += strings.Repeat(romans[i], r)
        }
    }

    return res
}
