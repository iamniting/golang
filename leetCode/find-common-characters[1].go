// https://leetcode.com/problems/find-common-characters
// Just sol to the problem, It does not include the I/O part

func commonChars(A []string) []string {

	calcCommon := func(s string, cmn []int) {
		CMN := make([]int, 26)

		for _, c := range s {
			CMN[rune(c)-'a'] += 1
		}

		for i := 0; i < 26; i++ {
			if cmn[i] != 0 && cmn[i] > CMN[i] {
				cmn[i] = CMN[i]
			}
		}
	}

	cmn := make([]int, 26)
	res := []string{}

	for _, c := range A[0] {
		cmn[rune(c)-'a'] += 1
	}

	for _, s := range A[1:] {
		calcCommon(s, cmn)
	}

	for i, n := range cmn {
		for j := 0; j < n; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}
