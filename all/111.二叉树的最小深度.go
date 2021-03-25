/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lm := minDepth(root.Left)
	rm := minDepth(root.Right)
	if lm == 0 || rm == 0 {
		return lm + rm + 1
	}

	return min(lm, rm) + 1

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

