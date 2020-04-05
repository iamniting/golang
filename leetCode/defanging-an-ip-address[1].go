// https://leetcode.com/problems/defanging-an-ip-address
// Just sol to the problem, It does not include the I/O part

func defangIPaddr(address string) string {

    for i:=0; i<len(address); i++ {

        if address[i] == '.' {

            address = address[:i] + "[.]" + address[i+1:]
            i += 2
        }
    }

    return address
}
