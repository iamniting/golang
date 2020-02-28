package main

import (
    "fmt"
    "math"
    "sort"
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

func GetMinMaxVertical(root * Node, hd int) (int, int) {
    if root == nil { return  hd+1, hd-1 }

    var min, max int

    min, max = GetMinMaxVertical(root.left, hd-1)
    min = int(math.Min(float64(min), float64(hd)))
    max = int(math.Max(float64(max), float64(hd)))

    max, max = GetMinMaxVertical(root.right, hd+1)
    min = int(math.Min(float64(min), float64(hd)))
    max = int(math.Max(float64(max), float64(hd)))

    return min, max
}

func VerticalOrderRecursive(
        root * Node, hd int, level int, m map[int]map[int][]int) {

    if root == nil { return }

    if _, ok := m[hd]; !ok {
        m[hd] = map[int][]int{}
    }
    m[hd][level] = append(m[hd][level], root.data)

    VerticalOrderRecursive(root.left, hd-1, level+1, m)
    VerticalOrderRecursive(root.right, hd+1, level+1, m)
}

func VerticalOrder(root * Node) {
    if root == nil { return }

    var m = map[int]map[int][]int{}

    VerticalOrderRecursive(root, 0, 1, m)

    min, max := GetMinMaxVertical(root, 0)

    // iterate the map in sorted fashion
    for ver:=min; ver<=max; ver++ {
        // get and sort the levels
        levels := []int{}
        for level := range m[ver] { levels = append(levels, level) }
        sort.Ints(levels)

        // iterate via sorted levels
        slice := []int{}
        for _, level := range levels {
            slice = append(slice, m[ver][level]...)
        }
        fmt.Println(slice)
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

    fmt.Println("\n\nVerticalOrder Traversal Recursively")
    VerticalOrder(root)
}
