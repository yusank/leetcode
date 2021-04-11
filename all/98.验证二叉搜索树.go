/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// var (
	// 	dfs func(*TreeNode,int,int) bool
	
	// )

	// // 思路 1： 给定一个范围 让每次遍历子节点的时候 缩小范围
	// dfs = func(node *TreeNode,up,low int) bool{
	// 	if node == nil {
	// 		return true
	// 	}

	// 	if node.Val > low && node.Val < up {
	// 		return dfs(node.Left, node.Val,low) && dfs(node.Right, up, node.Val)
	// 	}

	// 	return false
	// }
	
	// return dfs(root, math.MaxInt64,math.MinInt64)

	// 思路 2 中序遍历思路
	var (
		dfs2 func(*TreeNode) bool
		pre = math.MinInt64
	)
	
	dfs2 = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs2(node.Left) {
			return false
		}

		if node.Val <= pre {
			return false
		}

		pre = node.Val
		return dfs2(node.Right)
	}

	return dfs2(root)

}


// @lc code=end

