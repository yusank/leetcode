/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var (
		s1 = make([]int,0)
		s2 = make([]int,0)
		order func(*TreeNode,bool)
	)

	order = func(t *TreeNode, is bool) {
		if t == nil {
			if is {
				s1 = append(s1,-1)
			} else {
				s2=append(s2,-1)
			}
			
			return
		}

		if is {
			s1 = append(s1,t.Val)
		} else {
			s2=append(s2,t.Val)
		}

		

		order(t.Left,is)
		order(t.Right,is)
	}

	order(p,true)
	order(q,false)

	for i,v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true

}
// @lc code=end

