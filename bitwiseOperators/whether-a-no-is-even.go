package main

import "fmt"

func isNumberEven(n int) bool {
    // last bit will be always unset for even no
    return true && (n & 1 == 0)
}

func main() {
    for i := 0; i <= 16; i++ {
        fmt.Println(i, isNumberEven(i))
    }
}
