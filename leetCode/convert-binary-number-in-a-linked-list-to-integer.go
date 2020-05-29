package main

func getDecimalValue(head *ListNode) int {

	res := 0

	for head != nil {

		res = res << 1
		res = res | (head.Val & 1)

		head = head.Next
	}

	return res
}
