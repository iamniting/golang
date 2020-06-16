package main

func canWinNim(n int) bool {
	return (n-1)%4 == 0 || (n-2)%4 == 0 || (n-3)%4 == 0
}
