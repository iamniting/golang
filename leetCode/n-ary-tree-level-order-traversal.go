package main

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	iterate([]*Node{root}, &res)
	return res
}

func iterate(nodes []*Node, res *[][]int) {
	if len(nodes) == 0 {
		return
	}

	var slice []int
	var newNodes []*Node

	for _, node := range nodes {
		slice = append(slice, node.Val)
		newNodes = append(newNodes, node.Children...)
	}

	*res = append(*res, slice)
	iterate(newNodes, res)
}
