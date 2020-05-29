package main

func defangIPaddr(address string) string {

	res := ""

	for _, c := range address {

		if c == '.' {
			res += "[.]"
		} else {
			res += string(c)
		}
	}

	return res
}
