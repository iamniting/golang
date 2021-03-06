package main

func isPalindrome(head *ListNode) bool {

	slice := []int{}

	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	i, j := 0, len(slice)-1

	for i < j {
		if slice[i] != slice[j] {
			return false
		}
		i++
		j--
	}

	return true
}
