package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	type void struct{}
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
