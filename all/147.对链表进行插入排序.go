/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dum := &ListNode{Next:head}

	cur,last := head.Next, head 
	for cur != nil {
		if last.Val <= cur.Val {
			last = cur
		} else {
			prev := dum
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}

			last.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}

		cur = last.Next
	}

	return dum.Next
}
// @lc code=end

