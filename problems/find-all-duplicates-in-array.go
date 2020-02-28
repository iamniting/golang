// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package main

import (
    "fmt"
    "math"
)

func findDuplicates(nums []int) []int {
    res := []int{}

    for _, n := range nums {
        num := int(math.Abs(float64(n)))

        // if num is already occured add it in res slice
        if nums[num-1] < 0 {
            res = append(res, num)
        // if num is coming for first time mark it as visited with -num
        } else {
            nums[num-1] = -nums[num-1]
        }
    }

    return res
}

func main() {
    slice := []int{4, 3, 2, 7, 8, 2, 3, 1}

    res := findDuplicates(slice)

    fmt.Println(res)
}
