package main

import "sort"

func customSortString(S string, T string) string {
	magic := map[byte]int{}

	for i := range S {
		magic[S[i]] = i - 100
	}

	byteSlice := []byte(T)
	sort.Slice(byteSlice, func(i, j int) bool {
		return magic[byteSlice[i]] < magic[byteSlice[j]]
	})

	return string(byteSlice)
}
