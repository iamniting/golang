// https://www.youtube.com/watch?v=khTiTSZ5QZY

package main

func productExceptSelf(nums []int) []int {
	res := []int{1}
	left := 1
	right := 1
	l := len(nums)

	// get product of all left elements and store it in res slice
	for _, n := range nums[:l-1] {
		left *= n
		res = append(res, left)
	}

	// get product of right elements and store it in right var
	// change the val of res array by multiply left and right.
	for i := l - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right *= nums[i]
	}

	return res
}
