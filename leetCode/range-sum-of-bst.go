// https://leetcode.com/problems/range-sum-of-bst
// Just sol to the problem, It does not include the I/O part

func rangeSumBST(root *TreeNode, L int, R int) int {
    ans := 0
    PreOrderRecursive(root, L, R, &ans)
    return ans
}

func PreOrderRecursive(root * TreeNode, L, R int, ans *int) {
    if root == nil {
        return
    }
    if root.Val >= L && root.Val <= R {
        *ans += root.Val
    }
    PreOrderRecursive(root.Left, L, R, ans)
    PreOrderRecursive(root.Right, L, R, ans)
}
