package main

func isLongPressedName(name string, typed string) bool {

	i, j := 0, 0

	for i < len(name) || j < len(typed) {

		if i < len(name) && j < len(typed) && name[i] == typed[j] {
			i++
			j++
		} else if i > 0 && j < len(typed) && name[i-1] == typed[j] {
			j++
		} else {
			return false
		}
	}

	return true
}
