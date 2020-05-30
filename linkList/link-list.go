package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func InsertAtFirst(head *Node, data int) *Node {
	node := &Node{data, nil}
	node.next = head
	head = node
	return head
}

func InsertAtLast(head *Node, data int) *Node {
	if head == nil {
		node := &Node{data, nil}
		return node
	}
	// else
	var list *Node = head
	for list != nil {
		if list.next == nil {
			node := &Node{data, nil}
			list.next = node
			break
		}
		list = list.next
	}

	return head
}

func PrintList(head *Node) {
	list := head
	for list != nil {
		fmt.Print(list.data, " ")
		list = list.next
	}
}

func main() {
	// create first list
	var head1 *Node
	head1 = InsertAtFirst(head1, 1)
	head1 = InsertAtFirst(head1, 2)
	head1 = InsertAtFirst(head1, 3)
	head1 = InsertAtFirst(head1, 4)
	fmt.Println(head1)
	PrintList(head1)

	fmt.Println()
	fmt.Println()

	// create second list
	var head2 *Node
	head2 = InsertAtLast(head2, 1)
	head2 = InsertAtLast(head2, 2)
	head2 = InsertAtLast(head2, 3)
	head2 = InsertAtLast(head2, 4)
	fmt.Println(head2)
	PrintList(head2)
}
