package main

func longestCommonPrefix(strs []string) string {
	res := ""

	if len(strs) == 0 {
		return res
	}

	for i, s := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return res
			}

			if strs[j][i] != byte(s) {
				return res
			}
		}
		res += string(s)
	}
	return res
}
