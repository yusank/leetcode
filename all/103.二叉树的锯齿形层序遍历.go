/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var rs [][]int
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for l := 0; len(queue) > 0 ;l++ {
		vals := []int{}
		q := queue
		queue = nil
		for _,n := range q {
			vals =append(vals,n.Val)
			if n.Left != nil {
				queue =append(queue,n.Left)
			}

			if n.Right != nil {
				queue = append(queue,n.Right)
			}
		}

		if l %2 == 1 {
			for i,n := 0,len(vals);i < n/2;i++ {
				vals[i],vals[n-i-1] =vals[n-i-1],vals[i]
			}
		}

		rs=append(rs,vals)
	}

	return rs
}

/*
		 3
   		/ \
  	  9   20
	 / \  /  \
    6  4 15   7

*/

// @lc code=end

