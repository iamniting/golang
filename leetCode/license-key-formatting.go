package main

func licenseKeyFormatting(S string, K int) string {

	res, tmp := "", ""

	for i := len(S) - 1; i > -1; i-- {
		if S[i] == '-' {

		} else if S[i] > 96 {
			tmp = string(S[i]-32) + tmp
		} else {
			tmp = string(S[i]) + tmp
		}

		if len(tmp) == K || i == 0 && len(tmp) > 0 {
			if len(res) != 0 {
				res = "-" + res
			}
			res = tmp + res
			tmp = ""
		}
	}

	return res
}
