// https://leetcode.com/problems/binary-tree-right-side-view
// Just sol to the problem, It does not include the I/O part

func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{} }

    var queue []*TreeNode
    queue = append(queue, root, nil)

    res := []int{}

    for len(queue) > 0 {
        if queue[0] == nil {
            queue = queue[1:]
            if len(queue) > 0 { queue = append(queue, nil) }
            continue
        }

        parent := queue[0]
        queue = queue[1:]

        if len(queue) > 0 && queue[0] == nil {
            res = append(res, parent.Val)
        }

        if parent.Left != nil {
            queue = append(queue, parent.Left)
        }
        if parent.Right != nil {
            queue = append(queue, parent.Right)
        }
    }
    return res
}
