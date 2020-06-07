package main

func hasAlternatingBits(n int) bool {
	// opposite of last bit
	prev := (n & 1) ^ 1

	for ; n > 0; n >>= 1 {
		if n&1 == 1 && prev == 0 {
			prev = 1
			continue
		} else if n&1 == 0 && prev == 1 {
			prev = 0
			continue
		}
		return false
	}
	return true
}
