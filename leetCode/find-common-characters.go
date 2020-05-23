// https://leetcode.com/problems/find-common-characters
// Just sol to the problem, It does not include the I/O part

func commonChars(A []string) []string {
	res := []string{}

	for _, c := range A[0] {

		count := 0

		for i, str := range A[1:] {
			for j, char := range str {
				if c == char {
					count++
					A[i+1] = str[:j] + str[j+1:]
					break
				}
			}
		}

		if count == len(A[1:]) {
			res = append(res, string(c))
		}
	}
	return res
}
