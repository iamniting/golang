package main

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	recursive(root, 0, &res)
	return res
}

func recursive(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}

	if len(*res) <= level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)

	for _, node := range root.Children {
		recursive(node, level+1, res)
	}
}
