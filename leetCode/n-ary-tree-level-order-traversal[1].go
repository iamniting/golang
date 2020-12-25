package main

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var level int
	var res [][]int
	queue := []*Node{root, nil}

	for len(queue) != 0 {

		node := queue[0]
		queue = queue[1:]

		if node == nil && len(queue) != 0 {
			queue = append(queue, nil)
			level++
			continue
		} else if node == nil {
			// if last nil
			continue
		}

		queue = append(queue, node.Children...)

		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)
	}

	return res
}
