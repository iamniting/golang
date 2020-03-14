// https://leetcode.com/problems/add-binary
// Just sol to the problem, It does not include the I/O part

func addBinary(a string, b string) string {

    i, j := len(a)-1, len(b)-1
    res := ""
    carry := 0

    for i >= 0 || j >= 0 || carry != 0 {

        sum := carry

        if i >= 0 && a[i] == '1' {
            sum += 1
        }
        if j >= 0 && b[j] == '1' {
            sum += 1
        }

        if sum == 0 {
            res = "0" + res; carry = 0
        } else if sum == 1 {
            res = "1" + res; carry = 0
        } else if sum == 2 {
            res = "0" + res; carry = 1
        } else if sum == 3 {
            res = "1" + res; carry = 1
        }

        i--
        j--
    }
    return res
}
