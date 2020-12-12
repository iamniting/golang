package main

func interpret(command string) string {

	var res string

	for i := 0; i < len(command); {
		if command[i] == 'G' {
			res += "G"
			i++
		} else if command[i] == '(' && command[i+1] == ')' {
			res += "o"
			i += 2
		} else {
			res += "al"
			i += 4
		}
	}

	return res
}
