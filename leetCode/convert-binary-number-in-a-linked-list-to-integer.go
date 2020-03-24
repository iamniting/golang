// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer
// Just sol to the problem, It does not include the I/O part

func getDecimalValue(head *ListNode) int {

    res := 0

    for head != nil {

        res = res << 1
        res = res | (head.Val & 1)

        head = head.Next
    }

    return res
}
