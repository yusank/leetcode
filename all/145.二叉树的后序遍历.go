/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	var (
		nums  []int
		order func(*TreeNode)
	)

	order = func(n *TreeNode) {
		if n == nil {
			return
		}

		order(n.Left)
		order(n.Right)
		nums = append(nums, n.Val)
	}

	order(root)
	return nums
}

// @lc code=end

