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

func LevelOrderRecursive(root * Node, level int) {
    if root == nil { return }

    if level == 1 { fmt.Print(root.data, " "); return }

    LevelOrderRecursive(root.left, level - 1)
    LevelOrderRecursive(root.right, level - 1)
}

func LevelOrder(root * Node) {
    height := HeightRecursive(root)

    for h := 1; h <= height; h++ {
        LevelOrderRecursive(root, h)
        fmt.Println()
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

    fmt.Println("\n\nLevel By Level Traversal Recursively")
    LevelOrder(root)
}
