// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer
// Just sol to the problem, It does not include the I/O part

func subtractProductAndSum(n int) int {

    prod := 1
    sum := 0

    for n > 0 {
        last := n % 10
        n = n / 10

        prod *= last
        sum += last
    }

    return prod - sum
}
