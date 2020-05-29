package main

import "math"

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	max := 0
	for _, node := range root.Children {
		x := maxDepth(node)
		max = int(math.Max(float64(max), float64(x)))
	}

	return max + 1
}
