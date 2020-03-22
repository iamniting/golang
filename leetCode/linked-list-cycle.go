// https://leetcode.com/problems/linked-list-cycle
// Just sol to the problem, It does not include the I/O part

func hasCycle(head *ListNode) bool {

    if head == nil { return false }

    slow, fast := head, head.Next

    for fast != nil && fast.Next != nil {

        if slow == fast {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next
    }

    return false
}
