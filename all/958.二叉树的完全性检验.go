/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	stack := make([]*TreeNode,0)
	stack = append(stack,root)
	isEnd := false

	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]

		if isEnd && node != nil {
			return false
		}

		if node == nil {
			isEnd = true
			continue
		}

		stack = append(stack,node.Left)
		stack = append(stack,node.Right)
	}

	return true
}
// @lc code=end

