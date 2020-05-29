package main

func toLowerCase(str string) string {

	res := ""

	for _, c := range str {

		res += string(c | ' ')
	}

	return res
}
