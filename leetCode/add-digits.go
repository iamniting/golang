package main

func addDigits(num int) int {

	for num > 9 {
		i := num
		num = 0
		for ; i > 0; i /= 10 {
			num += i % 10
		}
	}
	return num
}
