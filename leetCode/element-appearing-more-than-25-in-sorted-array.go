package main

func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	count, length := 1, len(arr)

	for i := 1; i < length; i++ {
		if arr[i-1] == arr[i] {
			count++
		} else {
			count = 1
		}

		if count > length>>2 {
			return arr[i]
		}
	}

	return -1
}
