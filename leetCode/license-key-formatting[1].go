package main

func licenseKeyFormatting(S string, K int) string {

	var res string
	var count int

	for i := len(S) - 1; i > -1; i-- {
		if S[i] == '-' {
			continue
		} else if count == K {
			res = "-" + res
			count = 0
		}

		if S[i] > 96 {
			res = string(S[i]-32) + res
		} else {
			res = string(S[i]) + res
		}
		count++
	}

	return res
}
