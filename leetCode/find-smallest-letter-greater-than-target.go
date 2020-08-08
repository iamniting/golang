package main

func nextGreatestLetter(letters []byte, target byte) byte {

	for _, c := range letters {
		if c > target {
			return c
		}
	}

	return letters[0]
}
