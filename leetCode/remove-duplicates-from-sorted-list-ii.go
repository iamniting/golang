package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{-1, head}
	prev := dummy

	for head != nil {

		ptr := head.Next

		// run till elements are same
		for ptr != nil && head.Val == ptr.Val {
			ptr = ptr.Next
		}

		// if same element occurs
		if ptr != head.Next {
			prev.Next = ptr
			head = ptr
			continue
		}

		prev, head = head, head.Next
	}

	return dummy.Next
}
