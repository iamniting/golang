package main

import "fmt"


func MakeMaxHeap(slice []int, i int, size int) {
    if i > size/2 - 1 {
        fmt.Println(size, i)
        return
    }

    lchild := i * 2 + 1
    rchild := i * 2 + 2
    if slice[lchild] > slice[rchild] && slice[lchild] > slice[i] {
        slice[lchild], slice[i] = slice[i], slice[lchild]
        MakeMaxHeap(slice, lchild, size)
    } else if slice[rchild] > slice[lchild] && slice[rchild] > slice[i] {
        slice[rchild], slice[i] = slice[i], slice[rchild]
        MakeMaxHeap(slice, rchild, size)
    }
}

func MaxHeap(slice []int, size int) {
    for i:=size/2; i>=0; i-- {
        MakeMaxHeap(slice, i, size)
    }
}

func main() {
    heap := []int{5, 3, 17, 10, 84, 19, 6, 22, 9, 18, 20, 1, 24}
    size := len(heap)

    fmt.Println(heap)

    MaxHeap(heap, size)

    fmt.Println(heap)
}
