package main

func buddyStrings(A string, B string) bool {

	if len(A) != len(B) {
		return false
	}

	var count int
	a, b := [26]byte{}, [26]byte{}

	for i := range A {
		a[A[i]-'a']++
		b[B[i]-'a']++
		if A[i] != B[i] {
			count++
		}
	}

	if count != 0 && count != 2 {
		return false
	}

	for i := range a {
		if A == B && a[i] > 1 {
			return true
		}
		if a[i] != b[i] {
			return false
		}
	}
	return count == 2
}
