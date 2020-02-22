// https://leetcode.com/problems/trapping-rain-water
// https://www.youtube.com/watch?v=HmBbcDiJapY

package main

import (
    "fmt"
    "math"
)

func trap(h []int) int {
    res := 0

    for i:=0; i<len(h); i++ {
        leftMax, rightMax := h[i], h[i]

        // find max building from current building to left most building
        for j:=0; j<i; j++ {
            leftMax = int(math.Max(float64(leftMax), float64(h[j])))
        }

        // find max building from current building to right most building
        for j:=i+1; j<len(h); j++ {
            rightMax = int(math.Max(float64(rightMax), float64(h[j])))
        }

        // take minimun of leftMax and rightMax, 
        // that will be overflow value for current building
        overFlow := int(math.Min(float64(leftMax), float64(rightMax)))

        // substract the value of building itself from overFlow 
        // to get the waterCapacity of that building
        waterCapacity := overFlow - h[i]

        // add the waterCapacity of building in the res
        res += waterCapacity
    }
    return res
}

func main() {
    height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    res := trap(height)

    fmt.Println(res)
}
