package main

func sumOddLengthSubarrays(arr []int) int {
	/*
	   Consider the subarray that contains A[i],
	   we can take 0,1,2..,i elements on the left, we have i + 1 choices.
	   we can take 0,1,2..,len-1 elements on the right, we have len - i choices.

	   In total, there are (i + 1) * (len - i) subarrays, that contains A[i].
	   And there are ((i + 1) * (len - i) + 1) / 2 subarrays with odd length, that contains A[i].
	   A[i] will be counted ((i + 1) * (len - i) + 1) / 2 times.

	   Example of array [1,2,3,4,5]
	   1 X X X X subarray length 1
	   X 2 X X X subarray length 1
	   X X 3 X X subarray length 1
	   X X X 4 X subarray length 1
	   X X X X 5 subarray length 1
	   1 2 X X X subarray length 2
	   X 2 3 X X subarray length 2
	   X X 3 4 X subarray length 2
	   X X X 4 5 subarray length 2
	   1 2 3 X X subarray length 3
	   X 2 3 4 X subarray length 3
	   X X 3 4 5 subarray length 3
	   1 2 3 4 X subarray length 4
	   X 2 3 4 5 subarray length 4
	   1 2 3 4 5 subarray length 5

	   5 8 9 8 5 total times each index was added.
	   3 4 5 4 3 total times in odd length array with (x + 1) / 2
	   2 4 4 4 2 total times in even length array with x / 2
	*/
	var res int

	for index, num := range arr {
		res += num * (((index+1)*(len(arr)-index) + 1) / 2)
	}

	return res
}
