package main

func destCity(paths [][]string) string {

	m := make(map[string]int)

	for _, path := range paths {
		for _, p := range path {
			m[p]++
		}
	}

	for _, path := range paths {
		// start and destination city will be excatly one
		// so check for path[1] == 1
		if m[path[1]] == 1 {
			return path[1]
		}
	}
	return ""
}
