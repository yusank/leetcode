/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var (
		dfs func(*TreeNode) int
	)
	dfs =  func (node *TreeNode) int{
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		ans = max(ans,left+right+1)
		return max(left ,right)+1
	}

	ans=1
	dfs(root)

	return ans-1
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
// @lc code=end

