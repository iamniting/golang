// https://leetcode.com/problems/remove-outermost-parentheses
// Just sol to the problem, It does not include the I/O part

func removeOuterParentheses(S string) string {
    stack := []int{}
    res := ""

    for i, c := range S {
        if c == '(' {
            stack = append(stack, i)
        } else if c == ')' {
            start := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                res += S[start+1:i]
            }
        }
    }
    return res
}
