/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	var (
		nums  []int
		order func(*TreeNode)
	)

	order = func(n *TreeNode) {
		if n == nil {
			return
		}

		nums = append(nums, n.Val)
		order(n.Left)
		order(n.Right)
	}

	order(root)
	return nums
}

// @lc code=end

