package main

func findTheDifference(s string, t string) byte {

	var res rune

	for _, c := range s + t {
		res ^= c
	}

	return byte(res)
}
