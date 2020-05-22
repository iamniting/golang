// https://leetcode.com/problems/unique-email-addresses
// Just sol to the problem, It does not include the I/O part

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
