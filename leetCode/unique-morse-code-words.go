package main

func uniqueMorseRepresentations(words []string) int {

	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
		".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-",
		"..-", "...-", ".--", "-..-", "-.--", "--.."}

	set := make(map[string]bool)

	for _, str := range words {

		tmp := ""

		for _, c := range str {

			tmp += morse[c-'a']
		}

		set[tmp] = true
	}

	return len(set)
}
