// https://leetcode.com/problems/partition-list
// Just sol to the problem, It does not include the I/O part

func partition(head *ListNode, x int) *ListNode {

    lesserDummy := &ListNode{0, nil}
    greaterDummy := &ListNode{0, nil}

    lesser, greater := lesserDummy, greaterDummy

    // maintain two lists for lesser and greater or equal
    for head != nil {
        if head.Val < x {
            lesser.Next = head
            lesser = lesser.Next
        } else {
            greater.Next = head
            greater = greater.Next
        }
        head = head.Next
    }

    // combine those two lists
    lesser.Next = greaterDummy.Next
    greater.Next = nil

    return lesserDummy.Next
}
