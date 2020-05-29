package main

import "strings"

func numUniqueEmails(emails []string) int {

	m := make(map[string]bool)

	for _, email := range emails {
		strs := strings.Split(email, "@")
		strs[0] = strings.Split(strs[0], "+")[0]
		strs[0] = strings.ReplaceAll(strs[0], ".", "")
		m[strings.Join(strs, "@")] = true
	}

	return len(m)
}
