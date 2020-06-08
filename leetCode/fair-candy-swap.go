package main

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	set := map[int]bool{}

	for _, n := range A {
		sumA += n
		set[n] = true
	}

	for _, n := range B {
		sumB += n
	}

	diff := (sumA - sumB) / 2

	for _, n := range B {
		if set[n+diff] {
			return []int{n + diff, n}
		}
	}

	return nil
}
