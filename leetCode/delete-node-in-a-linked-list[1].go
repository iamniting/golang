// https://leetcode.com/problems/delete-node-in-a-linked-list
// Just sol to the problem, It does not include the I/O part

func deleteNode(node *ListNode) {

    *node = *node.Next

}