package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	p1, p2 := headA, headB
	c1, c2 := 0, 0

	// get len of list1
	for p1 != nil {
		c1++
		p1 = p1.Next
	}

	// get len of list2
	for p2 != nil {
		c2++
		p2 = p2.Next
	}

	// keep big list in p1
	if c1 > c2 {
		p1, p2 = headA, headB
	} else {
		p1, p2 = headB, headA
		c1, c2 = c2, c1
	}

	// get to the point where both list going to be same len
	for c1 > c2 {
		p1 = p1.Next
		c1--
	}

	// get the same node
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return nil
}
