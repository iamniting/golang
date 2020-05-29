package main

func removeNode(head *ListNode, n int) (*ListNode, int) {

	if head == nil {
		return nil, 1
	}

	var counter int
	head.Next, counter = removeNode(head.Next, n)

	if counter == n {
		return head.Next, counter + 1
	}

	return head, counter + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	head, _ = removeNode(head, n)

	return head
}
