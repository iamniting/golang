package main

func removeDuplicates(S string) string {

	flag := true

	for flag {
		flag = false

		for i := 1; i < len(S); i++ {
			if S[i-1] == S[i] {
				S = S[:i-1] + S[i+1:]
				flag = true
			}
		}
	}

	return S
}
