package main

func rotateString(A string, B string) bool {

	for i := 0; i < len(A); i++ {
		if A == B {
			return true
		}
		A = A[1:] + A[:1]
	}

	return A == B
}
