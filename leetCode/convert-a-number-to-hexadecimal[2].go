package main

func toHex(num int) string {

	if num == 0 {
		return "0"
	}

	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	var res string

	// 2's compliment of -ve int and keep +ve as it is
	num &= 0xffffffff

	for num != 0 {
		res = hex[num&15] + res
		num = num >> 4
	}

	return res
}
