package main

func numberOfLines(widths []int, S string) []int {
	lines, curNum := 1, 0

	for _, c := range S {
		curNum += widths[c-'a']

		if curNum > 100 {
			lines++
			curNum = widths[c-'a']
		}
	}
	return []int{lines, curNum}
}
