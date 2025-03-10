/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		lvl := []int{}
		n:= len(queue)
		f or i := 0; i < n; i++ {
			current := queue[0]
			queue = queue[1:] //pop

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			lvl = append(lvl, current.Val)


		res = append(res, lvl)

	}

	return res

  }
