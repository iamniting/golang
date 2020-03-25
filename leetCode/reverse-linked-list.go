// https://leetcode.com/problems/reverse-linked-list
// Just sol to the problem, It does not include the I/O part

func reverseList(head *ListNode) *ListNode {

    var prev *ListNode
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}
