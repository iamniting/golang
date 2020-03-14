// https://leetcode.com/problems/add-binary
// Just sol to the problem, It does not include the I/O part

func addBinary(a string, b string) string {

    i, j := len(a)-1, len(b)-1
    res := ""
    carry := 0

    for i >= 0 || j >= 0 || carry != 0 {
        sum := carry

        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }

        carry = sum / 2
        res = string(sum%2 + '0') + res
    }
    return res
}
