// https://leetcode.com/problems/rotate-list
// Just sol to the problem, It does not include the I/O part

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil { return head }

    var p * ListNode
    length := 1

    // get len and last node
    for p = head; p.Next != nil; p = p.Next {
        length++
    }

    if k % length == 0 { return head }

    // add head in last node
    p.Next = head

    // get len - k after removing loops from k
    length = length - (k % length)

    // get new last node
    p = head
    for ; length > 1; length-- {
        p = p.Next
    }

    // set nil in new last node
    head = p.Next
    p.Next = nil
    return head
}
