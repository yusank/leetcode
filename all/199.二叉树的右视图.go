/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 广度优先
	// var (
	// 	rs  []int
	// 	queue = make([]*TreeNode,0)
	// )

	// queue=append(queue,root)
	// for len(queue) != 0 {
	// 	size := len(queue)
	// 	for i := 0;i < size;i++ {
	// 		node := queue[0]
	// 		queue = queue[1:]
			
	// 		if node.Left != nil {
	// 			queue=append(queue,node.Left)
	// 		}

	// 		if node.Right != nil {
	// 			queue=append(queue,node.Right)
	// 		}

	// 		// 每一层最后一个元素
	// 		if i == size -1 {
	// 			rs = append(rs,node.Val)
	// 		}
	// 	}
	// }

	// 深度优先
	var (
		rs []int
		dfs func(*TreeNode,int)
	)

	dfs = func (node *TreeNode, dep int) {
		if node == nil {
			return
		}

		if dep == len(rs) {
			rs = append(rs, node.Val)
		}

		dep++
		dfs(node.Right,dep)
		dfs(node.Left, dep)
	}

	dfs(root,0)

	return rs
}
// @lc code=end

