package main

func minOperations(logs []string) int {

	var stack []string

	for _, ops := range logs {

		if ops == "../" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if ops != "./" {
			stack = append(stack, ops)
		}
	}

	return len(stack)
}
