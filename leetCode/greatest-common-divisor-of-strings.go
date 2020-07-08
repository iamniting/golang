package main

func gcdOfStrings(str1 string, str2 string) string {
	gcdOfInts := func(x, y int) int {
		if x < y {
			x, y = y, x
		}
		for x%y != 0 {
			x, y = y, x%y
		}
		return y
	}

	gcd := gcdOfInts(len(str1), len(str2))

	// check if str1 != str2
	for i := 0; i < gcd; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}

	// check if str1 is repeating itself again and again
	// eg. T + ... + T
	for i := range str1 {
		if str1[i] != str1[i%gcd] {
			return ""
		}
	}

	// check if str2 is repeating itself again and again
	// eg. T + ... + T
	for i := range str2 {
		if str2[i] != str2[i%gcd] {
			return ""
		}
	}

	return str1[:gcd]
}
