/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		rs [][]int
	)

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		ln := len(stack)
		vals := []int{}
		for i := 0; i < ln; i++ {
			node := stack[0]
			stack = stack[1:]
			vals = append(vals, node.Val)

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		rs = append([][]int{vals}, rs...)
	}

	return rs
}

// @lc code=end

