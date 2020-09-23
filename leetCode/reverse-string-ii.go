package main

func reverseStr(s string, k int) string {

	reverse := func(s []byte, i, j int) {
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}

	str := []byte(s)
	var flag bool

	for i := 0; i < len(s); i += k {

		flag = !flag

		if !flag {
			continue
		}

		if i+k > len(s) {
			reverse(str, i, len(s)-1)
			continue
		}
		reverse(str, i, i+k-1)
	}

	return string(str)
}
