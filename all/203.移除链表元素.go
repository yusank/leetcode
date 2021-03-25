/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	var (
		cur = head
	)

	for cur != nil {
		if cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}

		if cur.Next == nil || cur.Next.Val != val {
			cur = cur.Next
		}

	}

	return head
}

// @lc code=end

