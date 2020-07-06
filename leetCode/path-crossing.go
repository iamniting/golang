package main

func isPathCrossing(path string) bool {

	var ns, ew int
	set := map[[2]int]bool{{ns, ew}: true}

	for _, p := range path {

		if p == rune('N') {
			ns++
		} else if p == rune('S') {
			ns--
		} else if p == rune('E') {
			ew++
		} else if p == rune('W') {
			ew--
		}

		if set[[2]int{ns, ew}] {
			return true
		}

		set[[2]int{ns, ew}] = true
	}
	return false
}
