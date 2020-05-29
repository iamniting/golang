package main

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var PostOrder func(root *Node, res *[]int)

	PostOrder = func(root *Node, res *[]int) {

		for _, node := range root.Children {
			PostOrder(node, res)
		}
		*res = append(*res, root.Val)
	}

	res := []int{}
	PostOrder(root, &res)
	return res
}
