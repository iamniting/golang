// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
// Just sol to the problem, It does not include the I/O part

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
