// https://leetcode.com/problems/search-in-a-binary-search-tree
// Just sol to the problem, It does not include the I/O part

func searchBST(root *TreeNode, val int) *TreeNode {

    if root == nil { return nil }

    if root.Val == val { return root }

    var res *TreeNode

    if root.Val > val {
        res = searchBST(root.Left, val)
    } else if root.Val < val {
        res = searchBST(root.Right, val)
    }
    return res
}
