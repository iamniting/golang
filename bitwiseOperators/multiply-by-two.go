package main

import "fmt"

func multiplyByTwo(num int) int {
    /*
    32 = 0100000
    64 = 1000000
    32 << 1, which is 64
    */
    return num << 1
}

func main () {
    for i := 0; i <= 16; i++ {
        fmt.Println(i, multiplyByTwo(i))
    }
}
