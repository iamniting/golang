package main

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	res := make([]byte, 15)
	count, index := 0, 14

	for ; n > 0; n /= 10 {
		if count == 3 {
			res[index] = '.'
			count = 0
			index--
		}
		res[index] = byte(n%10 + 48)
		count++
		index--
	}

	return string(res[index+1:])
}
