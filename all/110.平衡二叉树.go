/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
        return true
    }
	 return abs(maxDepth(root.Left) - maxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -1*x
	}

	return x
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lc := maxDepth(root.Left)
	rc := maxDepth(root.Right)
	if lc > rc {
		return lc +1
	}

	return rc+1
}
// @lc code=end

