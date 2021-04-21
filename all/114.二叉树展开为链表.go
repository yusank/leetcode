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
	flattenOrg(root)
	return
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

// 原地
func flattenOrg(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			p := next
			for p.Right != nil {
				p = p.Right
			}

			p.Right = cur.Right
			cur.Left = nil
			cur.Right = next
		}

		cur = cur.Right
	}
}

// @lc code=end

