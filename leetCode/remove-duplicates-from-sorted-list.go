package main

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	ptr := head.Next
	prev := head

	for ptr != nil {

		if prev.Val == ptr.Val {
			prev.Next = ptr.Next
			ptr = ptr.Next
			continue
		}

		prev = ptr
		ptr = ptr.Next
	}

	return head
}
