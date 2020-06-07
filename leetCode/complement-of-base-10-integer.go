package main

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	one := 0
	// mark all bits 1 until i is 0
	for i := N; i > 0; i >>= 1 {
		one = one<<1 | 1
	}
	// xor the one and num e.g. 101 ^ 111 = 010
	return one ^ N
}
