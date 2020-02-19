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

func InsertNodeIterative(root * Node, data int) *Node {
    if root == nil {
        node := &Node{data, nil, nil}
        return node
    }

    node := root

    for node != nil {
        if node.data > data && node.left == nil {
            node.left = &Node{data, nil, nil}
            break
        } else if node.data < data && node.right == nil {
            node.right = &Node{data, nil, nil}
            break
        } else if node.data > data {
            node = node.left
        } else if node.data < data {
            node = node.right
        }
    }
    return root
}

func DeleteNodeIterative(root * Node, data int) *Node {
    node := root
    parent := root

    // find element node and parent
    for node != nil {
        if node.data == data {
            break
        } else if node.data > data {
            parent = node
            node = node.left
        } else if node.data < data {
            parent = node
            node = node.right
        }
    }

    if node == nil {
        return root
    }

    // if leaf node
    if node.left == nil && node.right == nil {
        if parent.left == node {
            parent.left = nil
        } else if parent.right == node {
            parent.right = nil
        } else if parent == node && root == node {
            return nil
        }
    // if only left child
    } else if node.left != nil && node.right == nil {
        if parent.left == node {
            parent.left = node.left
            node.left = nil
        } else if parent.right == node {
            parent.right = node.left
            node.left = nil
        } else if parent == node && root == node {
            node = node.left
            return node
        }
    // if only right child
    } else if node.left == nil && node.right != nil {
        if parent.left == node {
            parent.left = node.right
            node.right = nil
        } else if parent.right == node {
            parent.right = node.right
            node.right = nil
        } else if parent == node && root == node {
            node = node.right
            return node
        }
    // if both childs are present
    } else if node.left != nil && node.right != nil {
        // get right most child of left sub tree
        var successor * Node = node.left
        var succ_parent * Node = node

        for successor.right != nil {
            succ_parent = successor
            successor = successor.right
        }

        if succ_parent == node {
            node.data = successor.data
            node.left = successor.left
            successor = nil
        } else if succ_parent != node {
            succ_parent.right = successor.left
            successor.left = nil
            node.data = successor.data
        }
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
    slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}

    var root *Node
    for i, element := range slice {
        // call recursive and iterative method one by one
        if i & 1 == 1 {
            root = InsertNode(root, element)
        } else {
            root = InsertNodeIterative(root, element)
        }
    }

    fmt.Println("PreOrder")
    PreOrder(root)

    fmt.Println("\n\nInOrder")
    InOrder(root)
    fmt.Println("\n\nInOrder Iterative Method")
    InOrderIterative(root)

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

    fmt.Println("\nDelete")
    DeleteNodeIterative(root, 95)
    LevelOrderIterative(root)
    fmt.Println()
    DeleteNodeIterative(root, 90)
    LevelOrderIterative(root)
    fmt.Println()
    DeleteNodeIterative(root, 50)
    LevelOrderIterative(root)
}
