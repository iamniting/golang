package main

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	var res string
	var count int

	for ; n > 0; n /= 10 {
		if count == 3 {
			res = "." + res
			count = 0
		}
		res = string(n%10+48) + res
		count++
	}

	return res
}
