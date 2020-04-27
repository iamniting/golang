// https://leetcode.com/problems/fizz-buzz
// Just sol to the problem, It does not include the I/O part

func fizzBuzz(n int) []string {

    res := make([]string, n)

    for i:=1; i<=n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            res[i-1] = "FizzBuzz"
        } else if i % 3 == 0 {
            res[i-1] = "Fizz"
        } else if i % 5 == 0 {
            res[i-1] = "Buzz"
        } else {
            res[i-1] = string(strconv.Itoa(i))
        }
    }
    return res
}
