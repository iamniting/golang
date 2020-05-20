// https://leetcode.com/problems/maximum-depth-of-n-ary-tree
// Just sol to the problem, It does not include the I/O part

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
