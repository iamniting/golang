package main

func tribonacci(n int) int {
	if n < 2 {
		return n
	}

	t0, t1, t2 := 0, 1, 1

	for i := 3; i <= n; i++ {
		ti := t0 + t1 + t2
		t0, t1, t2 = t1, t2, ti
	}

	return t2
}
