/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumNode := &ListNode{
		Val:  0,
		Next: head,
	}

	var (
		first, second = dumNode, dumNode
	)

	for n != 0 {
		first = first.Next
		n--
	}

	for first.Next != nil {
		second = second.Next
		first = first.Next
	}

	second.Next = second.Next.Next
	return dumNode.Next
}

// @lc code=end

