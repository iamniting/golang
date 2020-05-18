// https://leetcode.com/problems/build-an-array-with-stack-operations
// Just sol to the problem, It does not include the I/O part

func buildArray(target []int, n int) []string {
	res := []string{}
	index := 0

	for i := 1; i <= n; i++ {
		if index >= len(target) {
			break
		}

		if target[index] == i {
			res = append(res, "Push")
			index++
		} else {
			res = append(res, []string{"Push", "Pop"}...)
		}
	}
	return res
}
