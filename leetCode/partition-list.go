// https://leetcode.com/problems/partition-list
// Just sol to the problem, It does not include the I/O part

func partition(head *ListNode, x int) *ListNode {
    if head == nil { return head }

    var dummy = &ListNode{-1, head}
    smallNode, prev, p := dummy, dummy, head

    // flag to keep track of greater or equal node
    flag := false

    for p != nil {
        // if node is greater or equal set the flag
        if p.Val >= x {
            flag = true
        // if greater or equal node is not occurred
        } else if !flag && p.Val < x {
            smallNode = p
        // if small node after the greater or equal node
        } else if flag && p.Val < x {
            prev.Next = p.Next
            p.Next = smallNode.Next
            smallNode.Next = p
            smallNode = smallNode.Next
            p = prev.Next
            continue
        }
        prev = p
        p = p.Next
    }

    return dummy.Next
}
