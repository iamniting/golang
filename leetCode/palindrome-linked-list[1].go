// https://leetcode.com/problems/palindrome-linked-list
// Just sol to the problem, It does not include the I/O part

func isPalindrome(head *ListNode) bool {
    // get middle node
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }

    if fast != nil { slow = slow.Next }

    // reverse the second half list
    firstHalf, secondHalf := head, reverseList(slow)

    // check for palindrome
    for firstHalf != nil && secondHalf != nil {
        if firstHalf.Val != secondHalf.Val {
            return false
        }
        firstHalf = firstHalf.Next
        secondHalf = secondHalf.Next
    }
    return true
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode

    for head != nil {
        head.Next, prev, head = prev, head, head.Next
    }
    return prev
}
