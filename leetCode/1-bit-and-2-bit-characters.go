package main

func isOneBitCharacter(bits []int) bool {

	var res bool
	i, l := 0, len(bits)

	for i < l {
		if bits[i] == 0 {
			i++
			res = true
		} else if bits[i+1] == 0 || bits[i+1] == 1 {
			i += 2
			res = false
		}
	}
	return res
}
