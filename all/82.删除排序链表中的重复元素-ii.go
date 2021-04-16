/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dum := &ListNode{Next: head}
	prev := dum

	// 双指针方案
	//  p指针在前 如果当前节点 与下一个节点的值相同 则p继续向前走 且same置位true
	// 当p遇到不一样的指针了 将 cur的下一个指向p的下一个 这样就可以减掉这范围内的所有相同的节点
	// 遇到不一样的节点时 p和 cur 均往前移动
	cur, p := dum, dum.Next
	same := false
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			same = true
			p = p.Next
		} else {
			if same {
				cur.Next = p.Next
				p = p.Next
				same = false
			} else {
				p = p.Next
				cur = cur.Next
			}
		}
	}

	if same {
		cur.Next = p.Next
	}

	return prev.Next
}

// @lc code=end

