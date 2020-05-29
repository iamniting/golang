package main

func findComplement(num int) int {

	one := 0
	// mark all bits 1 until i is 0
	for i := num; i > 0; i >>= 1 {
		one = one<<1 | 1
	}
	// xor the one and num e.g. 101 ^ 111 = 010
	return one ^ num
}
