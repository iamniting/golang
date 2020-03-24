// https://leetcode.com/problems/middle-of-the-linked-list
// Just sol to the problem, It does not include the I/O part

func middleNode(head *ListNode) *ListNode {

    slow, fast := head, head.Next

    for fast != nil && fast.Next != nil {

        slow = slow.Next
        fast = fast.Next.Next
    }

    // if len is odd
    if fast == nil { return slow }
    // if len is even
    return slow.Next
}
