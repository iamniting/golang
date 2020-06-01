package main

import "strconv"

func calPoints(ops []string) int {
	sum, n := 0, 0
	stack := make([]int, len(ops))

	for _, op := range ops {

		if op == "+" {
			n = stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, n)

		} else if op == "C" {
			n = -stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		} else if op == "D" {
			n = stack[len(stack)-1] * 2
			stack = append(stack, n)

		} else {
			n, _ = strconv.Atoi(op)
			stack = append(stack, n)
		}
		sum += n
	}

	return sum
}
