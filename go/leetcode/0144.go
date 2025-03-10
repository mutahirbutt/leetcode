/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}

	a := root.Val
	b := preorderTraversal(root.Left)
	c := preorderTraversal(root.Right)

	res = append(res, a)
	res = append(res, b...)
	res = append(res, c...)

	return res

}
