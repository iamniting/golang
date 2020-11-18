package main

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]byte, len(num1)+len(num2))

	for i := range res {
		res[i] = '0'
	}

	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {

			product := (num2[i] - '0') * (num1[j] - '0')
			sum := product + (res[i+j+1] - '0')
			res[i+j+1] = (sum % 10) + '0'
			res[i+j] = ((res[i+j] - '0') + sum/10) + '0'
		}
	}

	if res[0] == '0' {
		return string(res[1:])
	}

	return string(res)
}
