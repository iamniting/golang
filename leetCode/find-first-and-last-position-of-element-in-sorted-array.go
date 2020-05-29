package main

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	// get left most element
	res[0] = getLeft(nums, target)
	// return if element not found
	if res[0] == -1 {
		return res
	}
	// get right most element
	res[1] = getRight(nums, target)

	return res
}

func getLeft(nums []int, target int) int {
	res := -1
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] <= target {
			left = mid + 1
		}

		if nums[mid] == target {
			res = mid
		}
	}
	return res
}

func getRight(nums []int, target int) int {
	res := -1
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] >= target {
			right = mid - 1
		}

		if nums[mid] == target {
			res = mid
		}
	}
	return res
}
