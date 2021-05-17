/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		ln := len(stack)
		vals := []int{}
		for i := 0; i < ln; i++ {
			node := stack[0]
			stack = stack[1:]
			vals = append(vals, node.Val)

			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return res
}

// @lc code=end

