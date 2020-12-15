package main

func reformatNumber(number string) string {

	var res string

	for _, c := range number {
		if c == ' ' || c == '-' {
			continue
		}
		res += string(c)
	}

	number, res = res, ""
	lastGroup := len(number) % 3

	for i := 0; i < len(number); i++ {

		// if last group is of len 2
		if lastGroup == 2 && i+2 == len(number) {
			res += number[i:]
			break
			// if last group is of len 4
		} else if lastGroup == 1 && i+4 == len(number) {
			res += number[i:i+2] + "-" + number[i+2:]
			break
		} else {
			res += number[i : i+3]
			i += 2
			if i+1 != len(number) {
				res += "-"
			}
		}
	}

	return res
}
