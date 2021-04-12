/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	dHead := &ListNode{Val: -1, Next: head}
	prev := dHead
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	rNode := prev
	for i := 0; i < right-left+1; i++ {
		rNode = rNode.Next
	}

	lNode := prev.Next
	endNode := rNode.Next

	prev.Next = nil
	rNode.Next = nil
	reverseList(lNode)

	prev.Next = rNode
	lNode.Next = endNode

	return dHead.Next
}

func reverseList(head *ListNode) {
	if head == nil {
		return
	}

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
}

// @lc code=end

