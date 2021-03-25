/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// package main

// import (
// 	"log"
// )

// func main() {
// 	var (
// 		list = []int{1, 2, 2, 1}
// 		head = &ListNode{}
// 	)

// 	cur := head
// 	for i, v := range list {
// 		cur.Val = v
// 		if i != len(list)-1 {
// 			cur.Next = &ListNode{}
// 			cur = cur.Next
// 		}
// 	}

// 	isPalindrome(head)
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func (l *ListNode) Print() {
// 	c := l
// 	for c != nil {
// 		log.Println(c.Val)
// 		c = c.Next
// 	}
// }

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	mid := findMid(head)
	secondHalf := reverse(mid.Next)
	// secondHalf.Print()

	p1 := head
	p2 := secondHalf

	var rs = true
	for p2 != nil {
		if p1.Val != p2.Val {
			rs = false
			break
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	mid.Next = reverse(secondHalf)
	return rs
}

func findMid(head *ListNode) *ListNode {
	var (
		s, f = head, head
	)

	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	return s
}

func reverse(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

// @lc code=end
