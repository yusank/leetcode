/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		node1 = l1
		node2 = l2
		dHead = &ListNode{}
		last  int
	)

	cur := dHead
	for node1 != nil || node2 != nil {
		cur.Val += last
		last = 0
		if node1 != nil {
			cur.Val += node1.Val
			node1 = node1.Next
		}

		if node2 != nil {
			cur.Val += node2.Val
			node2 = node2.Next
		}

		if cur.Val >= 10 {
			last = 1
			cur.Val %= 10
		}

		if node1 != nil || node2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	if last > 0 {
		cur.Next = &ListNode{Val: 1}
	}

	return dHead
}

// @lc code=end

