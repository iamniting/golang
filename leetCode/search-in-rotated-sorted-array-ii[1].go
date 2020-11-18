package main

func search(nums []int, target int) bool {

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return true
		}

		// if we dont now which partition is sorted eg. 3,6,9,3,3,3,3 or 3,3,3,3,3,6,9
		if nums[start] == nums[mid] {
			start++
			// if first partition is sorted
		} else if nums[start] < nums[mid] {
			// if target exist in first partition
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
			// if second partition is sorted
		} else if nums[start] > nums[mid] {
			// if target exist in second partition
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
