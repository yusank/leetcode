/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	head := buildList([]int{1, 2, 3, 0})
	head.Print()
	result := rotateRight(head, 6)
	fmt.Println("result=")
	result.Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(arr []int) *ListNode {
	head := &ListNode{}
	p := head
	for i, v := range arr {
		p.Val = v
		if i+1 == len(arr) {
			break
		}
		next := &ListNode{}
		p.Next = next
		p = p.Next
	}

	return head
}

func (l *ListNode) Print() {
	c := l
	for c != nil {
		fmt.Print(c.Val)
		c = c.Next
		if c != nil {
			fmt.Print("->")
		}
	}
	fmt.Printf("\n")
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
loop:
	var (
		pre     = &ListNode{Next: head, Val: -1}
		p       = head
		ln  int = 0
	)

	for p != nil {
		ln++
		// fmt.Println(ln, k)
		// p.Print()
		if ln-k > 0 {
			pre = pre.Next
		}
		// pre.Print()
		if p.Next == nil {
			// stop here
			break
		}
		p = p.Next
	}

	if k == ln || k%ln == 0 {
		return head
	}

	if k > ln {
		// k > ln
		k = k % ln
		goto loop
	}

	result := pre.Next
	pre.Next = nil
	p.Next = head

	return result
}

// @lc code=end
