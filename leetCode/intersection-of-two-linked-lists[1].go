// https://leetcode.com/problems/intersection-of-two-linked-lists
// Just sol to the problem, It does not include the I/O part

func getIntersectionNode(headA, headB *ListNode) *ListNode {

    p1, p2 := headA, headB

    // run untill they did not get meet
    // if there is intersection they will meet at common node
    // if there is no intersection they will meet at nil
    for p1 != p2 {

        // p1 and p2 will run for len(list1) + len(list2)
        // which mean both p1 and p2 is going to end at same time
        // which mean they are going to reach intersected node at same time

        if p1 == nil {
            p1 = headB
        } else {
            p1 = p1.Next
        }

        if p2 == nil {
            p2 = headA
        } else {
            p2 = p2.Next
        }
    }

    return p1
}
