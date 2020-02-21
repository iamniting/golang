package main

import ( 
    "fmt"
    "math"
)

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

func HeightRecursive(root * Node) int {
    if root == nil { return 0 }

    return 1 + int(math.Max(
        float64(HeightRecursive(root.left)),
        float64(HeightRecursive(root.right))))
}

func LevelOrderIterative(root * Node) {
    if root == nil { return }

    height := HeightRecursive(root)

    var queue []*Node
    queue = append(queue, root)

    index := 0

    for len(queue) < int(math.Pow(2, float64(height))) - 1 {
        parent := queue[index]
        index++

        if parent != nil && parent.left != nil {
            queue = append(queue, parent.left)
        } else {
            queue = append(queue, nil)
        }

        if parent != nil && parent.right != nil {
            queue = append(queue, parent.right)
        } else {
            queue = append(queue, nil)
        }
    }

    for i, element := range queue {

        if element != nil {
            fmt.Print(element.data, " ")
        } else {
            fmt.Print("no", " ")
        }

        // if i + 2 is power of 2
        if (i + 2) & ((i + 2) - 1) == 0 {
            fmt.Println()
        }
    }
}

func main() {
    // first tree
    slice := []int{50, 20, 80, 10, 30, 60, 90, 5, 15, 25, 35, 55, 65, 85, 95}
    var root *Node

    fmt.Println("Insert Nodes Recursively")
    for _, element := range slice {
        fmt.Print(element, " ")
        root = InsertNodeRecursive(root, element)
    }

    fmt.Println("\n\nLevelOrder Traversal Iteratively")
    LevelOrderIterative(root)

    // second tree
    slice = []int{10, 20, 30, 40}
    root = nil

    fmt.Println("\nInsert Nodes Recursively")
    for _, element := range slice {
        fmt.Print(element, " ")
        root = InsertNodeRecursive(root, element)
    }

    fmt.Println("\n\nLevelOrder Traversal Iteratively")
    LevelOrderIterative(root)
}
