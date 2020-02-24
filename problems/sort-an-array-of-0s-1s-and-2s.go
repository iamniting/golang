// https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/
// Dutch National Flag Algorithm or 3-way Partitioning

package main

import "fmt"

func sortArray(arr []int) []int {
    low := 0
    mid := 0
    high := len(arr) - 1

    for mid <= high {
        // keep the track of 0s by low
        if arr[mid] == 0 {
            arr[low], arr[mid] = arr[mid], arr[low]
            low++
            mid++
        } else if arr[mid] == 1 {
            mid++
        // keep the track of 2s by high
        } else if arr[mid] == 2 {
            arr[mid], arr[high] = arr[high], arr[mid]
            high--
        }
    }
    return arr
}

func main() {
    slice := []int{2, 1, 0, 2, 2, 1, 1, 0, 1, 0, 0, 2, 1, 2, 0}
    fmt.Println(sortArray(slice))
}
