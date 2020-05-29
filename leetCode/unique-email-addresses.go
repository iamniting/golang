package main

import "strings"

func numUniqueEmails(emails []string) int {

	m := make(map[string]bool)

	for _, email := range emails {
		s := strings.Split(email, "@")
		email = ""

		for _, c := range s[0] {
			if c == '.' {
				continue
			} else if c == '+' {
				break
			}
			email += string(c)
		}
		email += "@" + s[1]
		m[email] = true
	}

	return len(m)
}
