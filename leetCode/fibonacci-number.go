// https://leetcode.com/problems/fibonacci-number
// Just sol to the problem, It does not include the I/O part

func fib(N int) int {

    a, b := 0, 1

    for i:=0; i<N; i++ {
        a, b = b, a + b
    }

    return a
}
