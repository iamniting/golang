// https://leetcode.com/problems/merge-two-sorted-lists
// Just sol to the problem, It does not include the I/O part

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

    var l * ListNode = &ListNode {0, nil }
    res := l

    for l1 != nil || l2 != nil {

        if l1 == nil {
            l.Next = l2
            break

        } else if l2 == nil {
            l.Next = l1
            break

        } else if l1.Val <= l2.Val {
            l.Next = l1
            l1 = l1.Next

        } else {
            l.Next = l2
            l2 = l2.Next

        }
        l = l.Next
    }

    return res.Next
}
