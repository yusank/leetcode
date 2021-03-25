/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	var (
		s, f = head, head
	)

	for f != nil {
		s = s.Next
		if f.Next == nil {
			return nil
		}

		f = f.Next.Next

		if f == s {
			p := head
			for p != s {
				p = p.Next
				s = s.Next
			}

			return p
		}
	}

	return nil
}

// @lc code=end

