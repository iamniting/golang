// https://leetcode.com/problems/valid-parentheses
// Just sol to the problem, It does not include the I/O part

func isValid(s string) bool {
    stack := []byte{}

    for _, c := range s {

        if c == '{' || c == '[' || c == '(' {
            stack = append(stack, byte(c))
            continue
        }

        if len(stack) == 0 {
            return false
        }

        top := stack[len(stack) - 1]

        if (c == '}' && top != '{') ||
           (c == ']' && top != '[') ||
           (c == ')' && top != '(') {
            return false
        }
        stack = stack[:len(stack) - 1]
    }

    return len(stack) == 0
}
