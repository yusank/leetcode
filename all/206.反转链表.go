/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "log"

func main() {
	var (
		list = []int{1, 2, 3, 4, 5}
		head = &ListNode{}
	)

	cur := head
	for _, v := range list {
		cur.Val = v
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	cur.Next = nil

	head = reverseList(head)
	head.Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	c := l
	for c != nil {
		log.Println(c.Val)
		c = c.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nh := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return nh
}

// @lc code=end
