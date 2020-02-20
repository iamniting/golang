// https://www.geeksforgeeks.org/count-pairs-from-two-arrays-having-sum-equal-to-k/

package main

import "fmt"

func CountPairs(a1 []int, a2 []int, k int) int {
    res := 0
    // create a set
    type void struct{}
    var value void
    set := make(map[int]void)

    // create a set of elements of slice 1
    for _, num := range a1 {
        set[num] = value
    }

    // check if (k - elements of slice 2) is present in set
    for _, num := range a2 {
        n := k - num
        if _, ok := set[n]; ok {
            res++
        }
    }

    return res

}

func main() {
    var a1 = []int{1, 1, 3, 4, 5, 6, 6}
    var a2 = []int{1, 4, 4, 5, 7}
    var k int = 10

    fmt.Println(CountPairs(a1, a2, k))
}
