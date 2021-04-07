/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		parent = make(map[int]*TreeNode)
		dfs func(*TreeNode)
		vis = make(map[int]bool)
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}

		if node.Right != nil {
			parent[node.Right.Val]= node
			dfs(node.Right)
		}
	}

	dfs(root)

	for p != nil {
		vis[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if vis[q.Val] {
			return q
		}

		q = parent[q.Val]
	}

	return nil
}
// @lc code=end

