package main

import (
	"fmt"

	avl "github.com/iamniting/golang/avlTree/pkg/avlTree"
)

func main() {
	slice := []int{20, 30, 40, 35, 38, 37, 10}
	var root *avl.Node

	fmt.Println("Insert Nodes in AVL Tree Recursively")
	for _, element := range slice {
		fmt.Print(element, " ")
		root = avl.InsertInAvl(root, element)
	}

	fmt.Println("\n\nLevelOrder Traversal Recursively")
	avl.LevelOrder(root)

	fmt.Println("\nDelete Node in AVL Tree Recursively")
	root = avl.DeleteInAvl(root, 35)

	fmt.Println("\nLevelOrder Traversal Recursively")
	avl.LevelOrder(root)
}
