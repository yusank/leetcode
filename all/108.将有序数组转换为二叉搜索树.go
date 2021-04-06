/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left,right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left+right)/2
	r := &TreeNode{Val: nums[mid]}
	r.Left = build(nums,left,mid-1)
	r.Right = build(nums,mid+1,right)

	return r
}
// @lc code=end

