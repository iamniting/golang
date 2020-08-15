package main

func licenseKeyFormatting(S string, K int) string {

	res := make([]byte, len(S)+len(S)/K)
	idx := len(res) - 1
	var count int

	for i := len(S) - 1; i > -1; i-- {
		if S[i] == '-' {
			continue
		} else if count == K {
			res[idx] = '-'
			idx--
			count = 0
		}

		if S[i] > 96 {
			res[idx] = S[i] - 32
		} else {
			res[idx] = S[i]
		}
		count++
		idx--
	}

	return string(res[idx+1:])
}
