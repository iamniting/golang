package main

import "fmt"

func sortArray(arr []int) {
    low := 0
    high := len(arr) - 1

    for low <= high {

        for arr[low] == 0 { low++ }

        for arr[high] == 1 { high-- }

        if low < high {
            arr[low], arr[high] = arr[high], arr[low]
        }
    }
}

func main() {
    slice := []int{0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0}
    sortArray(slice)
    fmt.Println(slice)
}

