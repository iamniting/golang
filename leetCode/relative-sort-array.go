// https://leetcode.com/problems/relative-sort-array
// Just sol to the problem, It does not include the I/O part

func relativeSortArray(arr1 []int, arr2 []int) []int {

	bucket := make([]int, 1001)

	for _, n := range arr1 {
		bucket[n]++
	}

	index := 0
	for _, n := range arr2 {
		for bucket[n] > 0 {
			arr1[index] = n
			index++
			bucket[n]--
		}
	}

	for i, c := range bucket {
		for c > 0 {
			arr1[index] = i
			index++
			c--
		}
	}

	return arr1
}
