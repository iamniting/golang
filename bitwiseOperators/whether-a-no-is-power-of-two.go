package main

import "fmt"

func isPowerOfTwo(n int) bool {
    /*
    if n is power of 2, there will be only one bit set
    and result of n & (n-1) will be always zero e.g.
    64 = 1000000
    63 = 0111111
    64 & 63 will be zero
    */
    return (n != 0) && (n & (n - 1) == 0)
}

func main() {
    for i := 0; i <= 16; i++ {
        fmt.Println(i, isPowerOfTwo(i))
    }
}
