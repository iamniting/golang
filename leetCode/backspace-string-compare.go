package main

func backspaceCompare(S string, T string) bool {

	var stack1, stack2 string

	for _, c := range S {
		if c == '#' && len(stack1) > 0 {
			stack1 = stack1[0 : len(stack1)-1]
		} else if c != '#' {
			stack1 += string(c)
		}
	}

	for _, c := range T {
		if c == '#' && len(stack2) > 0 {
			stack2 = stack2[0 : len(stack2)-1]
		} else if c != '#' {
			stack2 += string(c)
		}
	}

	return stack1 == stack2
}
