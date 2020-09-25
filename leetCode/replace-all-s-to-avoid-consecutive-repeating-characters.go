package main

func modifyString(s string) string {

	res := []rune(s)

	for i := range res {
		if res[i] != '?' {
			continue
		}

		for _, c := range "abc" {
			if (i == 0 || res[i-1] != c) && (i == len(res)-1 || res[i+1] != c) {
				res[i] = c
			}
		}
	}

	return string(res)
}
