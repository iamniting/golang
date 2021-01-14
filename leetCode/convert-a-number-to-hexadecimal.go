package main

func toHex(num int) string {

	if num == 0 {
		return "0"
	}

	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	var res string

	// 2's compliment of negative num
	if num < 0 {
		num = (1 << 32) + num
	}

	for num != 0 {
		res = hex[num%16] + res
		num = num / 16
	}

	return res
}
