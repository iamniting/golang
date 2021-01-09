package main

func totalMoney(n int) int {

	var res int
	weeks := n / 7
	days := n % 7

	for i, money := 0, 28; i < weeks; i, money = i+1, money+7 {
		res += money
	}

	for i, money := 0, weeks+1; i < days; i, money = i+1, money+1 {
		res += money
	}

	return res
}
