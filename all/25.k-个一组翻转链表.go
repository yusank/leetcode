/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	dum := &ListNode{Next: head}
}

func reverse(head *ListNode) *ListNode {
	var (
		prev *ListNode
		cur  = head
	)

	for cur != nil {
		tempNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = tempNext
	}

	return prev
}

// @lc code=end

