/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		path   []int
		result [][]int
		dfs    func(node *TreeNode, left int)
	)

	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}

		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		dfs(node.Left, left)
		dfs(node.Right, left)
	}

	dfs(root, targetSum)
	return result
}

// @lc code=end

