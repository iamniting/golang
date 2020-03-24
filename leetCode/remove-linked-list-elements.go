// https://leetcode.com/problems/remove-linked-list-elements
// Just sol to the problem, It does not include the I/O part

func removeElements(head *ListNode, val int) *ListNode {

    p := head
    var prev *ListNode = nil

    for p != nil {

        // if head node
        if prev == nil && p.Val == val {
            head = p.Next
            p = p.Next
        // if non-head node
        } else if p.Val == val {
            prev.Next = p.Next
            p = p.Next
        } else {
            prev = p
            p = p.Next
        }
    }

    return head
}
