/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 思路
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow,fast := head,head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 找到中间 然后链表断开
	temp := slow.Next
	slow.Next = nil

	// 左右部分继续拆  最终个返回一个节点
	left := sortList(head)
	right := sortList(temp)

	h := &ListNode{}
	ans := h
	// 一层一层对比合并
	//  最开始是left right 各位一个或 0 个节点的链表 对比合并成 2 个节点链表 返回到上一层
	// 	上一层拿到两个节点的 left right 然后再对比合并 一直到第一层
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}

		h = h.Next
	}

	if left != nil {
		h.Next = left
	}

	if right != nil {
		h.Next = right
	}

	return ans.Next
}
// @lc code=end

