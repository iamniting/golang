// https://leetcode.com/problems/n-ary-tree-preorder-traversal
// Just sol to the problem, It does not include the I/O part

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var PreOrder func(root *Node, res *[]int)

	PreOrder = func(root *Node, res *[]int) {

		*res = append(*res, root.Val)
		for _, node := range root.Children {
			PreOrder(node, res)
		}
	}

	res := []int{}
	PreOrder(root, &res)
	return res
}
