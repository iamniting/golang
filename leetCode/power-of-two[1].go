package main

func isPowerOfTwo(n int) bool {

	var bits int

	for ; n > 0; n = n >> 1 {
		bits += n & 1
	}

	return bits == 1
}
