// https://www.youtube.com/watch?v=HmBbcDiJapY

package main

import "math"

func trap(height []int) int {
	res := 0

	for i := 0; i < len(height); i++ {
		leftMax, rightMax := height[i], height[i]

		// find max building from current building to left most building
		for j := 0; j < i; j++ {
			leftMax = int(math.Max(float64(leftMax), float64(height[j])))
		}

		// find max building from current building to right most building
		for j := i + 1; j < len(height); j++ {
			rightMax = int(math.Max(float64(rightMax), float64(height[j])))
		}

		// take minimun of leftMax and rightMax,
		// that will be overflow value for current building
		overFlow := int(math.Min(float64(leftMax), float64(rightMax)))

		// subtract the value of building itself from overFlow
		// to get the waterCapacity of that building
		waterCapacity := overFlow - height[i]

		// add the waterCapacity of building in the res
		res += waterCapacity
	}
	return res
}
