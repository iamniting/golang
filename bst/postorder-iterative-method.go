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

func PostOrderIterative(root * Node) {
    if root == nil {
        return
    }
    var stack1 = []*Node{root}
    var stack2 []*Node
    var node *Node = root
    for len(stack1) > 0 {
        node = stack1[len(stack1) - 1]
        stack1 = stack1[0:len(stack1) - 1]
        stack2 = append(stack2, node)
        if node.left != nil {
            stack1 = append(stack1, node.left)
        }
        if node.right != nil {
            stack1 = append(stack1, node.right)
        }
    }

    for i:=len(stack2)-1; i>=0; i-- {
        fmt.Print(stack2[i].data, " " )
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

    fmt.Println("\n\nPostOrder Traversal Iteratively")
    PostOrderIterative(root)
}
