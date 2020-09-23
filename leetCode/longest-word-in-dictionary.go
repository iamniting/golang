package main

func longestWord(words []string) string {

	var res string
	set := map[string]interface{}{}

	for _, word := range words {
		set[word] = nil
	}

	for _, word := range words {
		// if we can not change the ans then there is no need to check the set
		if !(len(word) > len(res) || len(word) == len(res) && word < res) {
			continue
		}

		var flag bool

		// check all prefixes
		for i := range word {
			if _, ok := set[word[:i+1]]; !ok {
				flag = true
				break
			}
		}

		// if all prefixes found then change the res
		if !flag {
			res = word
		}
	}

	return res
}
