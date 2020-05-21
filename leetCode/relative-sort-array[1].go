// https://leetcode.com/problems/relative-sort-array
// Just sol to the problem, It does not include the I/O part

func relativeSortArray(arr1 []int, arr2 []int) []int {

	index := 0

	for i := range arr2 {
		for j := index; j < len(arr1); j++ {
			if arr1[j] == arr2[i] {
				arr1[index], arr1[j] = arr1[j], arr1[index]
				index++
			}
		}
	}

	slice := arr1[index:]
	sort.Ints(slice)
	return arr1
}
