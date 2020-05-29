package main

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var first *ListNode = head
	var second *ListNode = head.Next
	var prev *ListNode = nil

	head = second

	for first != nil && first.Next != nil {
		second = first.Next
		third := first.Next.Next

		second.Next = first
		first.Next = third

		if prev != nil {
			prev.Next = second
		}

		prev = first
		first = first.Next
	}

	return head
}
