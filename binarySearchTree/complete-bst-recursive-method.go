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
    if level == 0 { return }

    if level == 1 && root != nil { fmt.Print(root.data, " "); return }

    if level == 1 && root == nil { fmt.Print("no", " "); return }

    if root != nil && root.left != nil {
        LevelOrderRecursive(root.left, level - 1)
    } else {
        LevelOrderRecursive(nil, level -1 )
    }
    if root != nil && root.right != nil {
        LevelOrderRecursive(root.right, level - 1)
    } else {
        LevelOrderRecursive(nil, level - 1)
    }
}

func LevelOrder(root * Node) {
    height := HeightRecursive(root)

    for h := 1; h <= height; h++ {
        LevelOrderRecursive(root, h)
        fmt.Println()
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

    fmt.Println("\n\nLevelOrder Traversal Recursively")
    LevelOrder(root)


    // second tree
    slice = []int{10, 20, 30, 40}
    root = nil

    fmt.Println("\nInsert Nodes Recursively")
    for _, element := range slice {
        fmt.Print(element, " ")
        root = InsertNodeRecursive(root, element)
    }

    fmt.Println("\n\nLevelOrder Traversal Recursively")
    LevelOrder(root)
}
