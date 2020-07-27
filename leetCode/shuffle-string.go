package main

func restoreString(s string, indices []int) string {

	res := make([]byte, len(s))

	for i, n := range indices {
		res[n] = s[i]
	}

	return string(res)
}
