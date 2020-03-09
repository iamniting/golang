// https://leetcode.com/problems/add-two-numbers
// Just sol to the problem, It does not include the I/O part

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var res *ListNode = nil
    p, p1, p2 := res, l1, l2
    sum, carry := 0, 0

    for p1 != nil || p2 != nil || carry != 0 {

        sum = carry

        if p1 != nil {
            sum += p1.Val
            p1 = p1.Next
        }
        if p2 != nil {
            sum += p2.Val
            p2 = p2.Next
        }

        carry = sum / 10

        if res == nil {
            res = &ListNode{sum%10, nil}
            p = res
        } else {
            p.Next = &ListNode{sum%10, nil}
            p = p.Next
        }
    }
    return res
}
