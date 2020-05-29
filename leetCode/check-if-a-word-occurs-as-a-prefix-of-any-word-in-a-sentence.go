package main

func isPrefixOfWord(sentence string, searchWord string) int {

	wCount := 1

	for i := range sentence {
		if sentence[i] == ' ' {
			wCount++
		}

		if i == 0 || sentence[i-1] == ' ' {
			j := 0

			for j < len(searchWord) && i < len(sentence) && sentence[i] == searchWord[j] {
				i++
				j++
			}

			if j == len(searchWord) {
				return wCount
			}
		}
	}

	return -1
}
