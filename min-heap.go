package main

import "fmt"


func MakeMinHeap(slice []int, i int, size int) {
    if i > size/2 - 1 {
        return
    }

    lchild := i * 2 + 1
    rchild := i * 2 + 2

    if slice[lchild] < slice[rchild] && slice[lchild] < slice[i] {
        slice[lchild], slice[i] = slice[i], slice[lchild]
        MakeMinHeap(slice, lchild, size)
    } else if slice[rchild] < slice[lchild] && slice[rchild] < slice[i] {
        slice[rchild], slice[i] = slice[i], slice[rchild]
        MakeMinHeap(slice, rchild, size)
    }
}

func MinHeap(slice []int, size int) {
    for i:=size/2; i>=0; i-- {
        MakeMinHeap(slice, i, size)
    }
}

func main() {
    heap := []int{5, 3, 17, 10, 84, 19, 6, 22, 9, 18, 20, 1, 24}
    size := len(heap)

    fmt.Println(heap)

    MinHeap(heap, size)

    fmt.Println(heap)
}
