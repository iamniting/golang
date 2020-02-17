package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func InsertNode(root *Node, data int) *Node {
    if root == nil {
        node := &Node{data, nil, nil}
        return node
    } else if  data < root.data {
        root.left = InsertNode(root.left, data)
    } else if data > root.data {
        root.right = InsertNode(root.right, data)
    }
    return root
}

func PreOrder(root * Node) {
    if root == nil {
        return
    }
    fmt.Print(root.data, " ")
    PreOrder(root.left)
    PreOrder(root.right)
}

func InOrder(root * Node) {
    if root == nil {
        return
    }
    InOrder(root.left)
    fmt.Print(root.data, " ")
    InOrder(root.right)
}

func PostOrder(root * Node) {
    if root == nil {
        return
    }
    PostOrder(root.left)
    PostOrder(root.right)
    fmt.Print(root.data, " ")
}

func main() {
    var root *Node
    root = InsertNode(root, 5)
    root = InsertNode(root, 3)
    root = InsertNode(root, 2)
    root = InsertNode(root, 4)
    root = InsertNode(root, 7)
    root = InsertNode(root, 6)
    root = InsertNode(root, 8)

    PreOrder(root)
    fmt.Println()
    InOrder(root)
    fmt.Println()
    PostOrder(root)
}
