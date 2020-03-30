// https://leetcode.com/problems/linked-list-cycle-ii
// Just sol to the problem, It does not include the I/O part

func detectCycle(head *ListNode) *ListNode {
    if head == nil { return nil }

    type void struct {}
    m := make(map[*ListNode]void)

    for head != nil {
        if _, ok := m[head]; ok {
            return head
        }
        m[head] = void{}
        head = head.Next
    }

    return nil
}
