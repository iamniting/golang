package main

func halvesAreAlike(s string) bool {

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	var a, b int

	for i, c := range s {
		if vowels[c] {
			if i < len(s)/2 {
				a++
			} else {
				b++
			}
		}
	}

	return a == b
}
