package main

func numSpecialEquivGroups(A []string) int {
	set := make(map[[52]int8]interface{})

	for _, s := range A {

		key := [52]int8{}

		for i, c := range s {

			if i&1 == 1 {
				key[c-'a']++
			} else {
				key[c-'a'+26]++
			}
		}
		set[key] = nil
	}

	return len(set)
}
