// https://leetcode.com/problems/longest-valid-parentheses/
// https://www.youtube.com/watch?v=r0-zx5ejdq0

package main

import (
    "fmt"
    "math"
)

func getMaxDiff(slice []int, l int) int {
    max := 0

    slice = append(slice, l+1)
    slice = append([]int{0}, slice...)

    for i:=0; i<len(slice)-1; i++ {
        max = int(math.Max(float64(max), float64(slice[i+1] - slice[i] - 1)))
    }
    return max
}

func longestValidParentheses(str string) int {
    var stack = []int{}

    for i, s := range str {
        l := len(stack)

        if s == '(' {
            stack = append(stack, i+1)
        } else if s == ')' && len(stack) > 0 && str[stack[l-1] - 1] == '(' {
            stack = stack[:len(stack) - 1]
        } else if s == ')' && len(stack) == 0 || str[stack[l-1] - 1] == ')' {
            stack = append(stack, i+1)
        }
    }

    if len(stack) == 0 { return len(str) }

    return getMaxDiff(stack, len(str))
}

func main() {
    var str string
    str = "(()"
    fmt.Println(longestValidParentheses(str))
    str = ")()())"
    fmt.Println(longestValidParentheses(str))
}
