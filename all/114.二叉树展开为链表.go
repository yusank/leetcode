/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		order func(node *TreeNode)
		nt    = &TreeNode{}
		cur   = nt
	)

	order = func(node *TreeNode) {
		if node == nil {
			return
		}

		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right

		order(node.Left)
		order(node.Right)
	}

	order(root.Left)
	order(root.Right)

	root.Left = nil
	root.Right = nt.Right
}

// @lc code=end

