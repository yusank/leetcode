/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	var (
		ans []int
		stack = make([]*TreeNode,0)
	)

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans,root.Val)
		root = root.Right
	}

	
	return ans
}
// @lc code=end

