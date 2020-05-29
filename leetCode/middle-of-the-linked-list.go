package main

func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
	}

	// if len is odd
	if fast == nil {
		return slow
	}
	// if len is even
	return slow.Next
}
