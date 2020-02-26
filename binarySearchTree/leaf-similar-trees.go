// https://leetcode.com/problems/leaf-similar-trees/
// Just sol to the problem, It does not include the I/O part

func inorder(root *TreeNode, slice *[]int) {
    if root == nil { return }

    inorder(root.Left, slice)
    if root.Left == nil && root.Right == nil {
        *slice = append(*slice, root.Val)
    }
    inorder(root.Right, slice)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

    if root1 == nil || root2 == nil { return false }

    slice1 := []int{}
    slice2 := []int{}

    inorder(root1, &slice1)
    inorder(root2, &slice2)

    if len(slice1) != len(slice2) { return false }

    for i, val := range slice1 {
        if val != slice2[i] { return false }
    }
    return true
}
