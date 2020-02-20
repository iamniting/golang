package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func InsertNodeRecursive(root *Node, data int) *Node {
    if root == nil {
        node := &Node{data, nil, nil}
        return node
    } else if  data < root.data {
        root.left = InsertNodeRecursive(root.left, data)
    } else if data > root.data {
        root.right = InsertNodeRecursive(root.right, data)
    }
    return root
}

func InOrderIterative(root * Node) {
    if root == nil {
        return
    }
    var stack []*Node
    var node *Node = root

    for node != nil || len(stack) > 0 {
        if node != nil {
            stack = append(stack, node)
            node = node.left
        } else if node == nil {
            node = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            fmt.Print(node.data, " ")
            node = node.right
        }
    }
}

func main() {
    slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}

    var root *Node

    fmt.Println("Insert Nodes Recursively")
    for _, element := range slice {
        fmt.Print(element, " ")
        root = InsertNodeRecursive(root, element)
    }

    fmt.Println("\n\nInOrder Traversal Iteratively")
    InOrderIterative(root)
}
