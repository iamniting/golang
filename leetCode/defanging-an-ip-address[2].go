// https://leetcode.com/problems/defanging-an-ip-address
// Just sol to the problem, It does not include the I/O part

func defangIPaddr(address string) string {

    return strings.Replace(address, ".", "[.]", -1)
}
