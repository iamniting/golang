package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	var prev *ListNode
	l1, l2 := list1, list2

	for l1 != nil {
		if l1.Val == a {
			prev.Next = l2
		}

		prev = l1
		l1 = l1.Next

		if prev.Val == b {
			break
		}
	}

	for l2 != nil {
		if l2.Next == nil {
			l2.Next = l1
			break
		}
		l2 = l2.Next
	}

	return list1
}
