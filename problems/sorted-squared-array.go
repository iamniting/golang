// https://www.geeksforgeeks.org/sort-array-converting-elements-squares/
// https://leetcode.com/problems/squares-of-a-sorted-array/submissions/
// https://www.youtube.com/watch?v=4eWKHLSRHPY

package main

import (
    "fmt"
    "math"
)

func sortedSquaredArray(a []int) []int {
    // left most index
    i := 0
    // right most index
    j := len(a) - 1
    // result array right most index
    length := j
    // result array
    res := make([]int, j+1)

    // iterate both the pointer untill they meet
    for i <= j {
        // if left index is greater than right, then add a[i] in res
        if math.Abs(float64(a[i])) > math.Abs(float64(a[j])) {
            res[length] = a[i] * a[i]
            i++
        // if right index is greater then left, then add a[j] in res
        } else {
            res[length] = a[j] * a[j]
            j--
        }
        length--
    }

    return res
}

func main() {
    arr := []int{-6, -4, 1, 2, 3, 4, 5}
    fmt.Println(sortedSquaredArray(arr))
    arr = []int{-7,-3,2,3,11}
    fmt.Println(sortedSquaredArray(arr))
}
