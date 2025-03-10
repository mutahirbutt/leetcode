package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/bfs solution
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	q := []*TreeNode{}
	q = append(q, root)
	depth := 0

	for len(q) != 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}

		}
		depth++
	}
	return depth
}
