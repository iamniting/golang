package main

import (
    "fmt"
)

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

func LevelOrderIterative(root * Node) {
    if root == nil {
        return
    }
    var queue []*Node
    queue = append(queue, root)

    for len(queue) > 0 {
        parent := queue[0]
        queue = queue[1:]

        fmt.Print(parent.data, " ")
        if parent.left != nil {
            queue = append(queue, parent.left)
        }
        if parent.right != nil {
            queue = append(queue, parent.right)
        }
    }
}

func Search(root * Node, data int) bool {
    if root == nil {
        return false
    } else if root.data == data {
        return true
    } else if root.data > data {
        return Search(root.left, data)
    } else if root.data < data {
        return Search(root.right, data)
    }
    return false
}

func SearchIterative(root * Node, data int) {
    if root == nil {
        fmt.Println("Tree is empty")
        return
    }

    var node * Node = root
    for node != nil {
        if data == node.data {
            fmt.Println("Found element", data)
            return
        } else if data < node.data {
            node = node.left
        } else if data > node.data {
            node = node.right
        }
    }
    fmt.Println("Not found element", data)
}

func main() {
    var root *Node
    slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 96}

    for _, element := range slice {
        root = InsertNode(root, element)
    }

    fmt.Println("PreOrder")
    PreOrder(root)

    fmt.Println("\n\nInOrder")
    InOrder(root)

    fmt.Println("\n\nPostOrder")
    PostOrder(root)

    fmt.Println("\n\nLevelOrder")
    LevelOrderIterative(root)

    fmt.Println("\n\nSearch")
    res := Search(root, 5)
    fmt.Println(5, res)
    res = Search(root, 50)
    fmt.Println(50, res)
    res = Search(root, 100)
    fmt.Println(100, res)

    fmt.Println("\nSearch Iterative")
    SearchIterative(root, 5)
    SearchIterative(root, 50)
    SearchIterative(root, 100)
}
