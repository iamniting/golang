package main

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
