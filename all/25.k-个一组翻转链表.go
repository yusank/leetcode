/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */


 package main

 import "fmt"

 func main() {
	head := buildList([]int{1,2,3,4,5})
	head.print()
	reverseKGroup(head,3).print()
 }

type ListNode struct {
     Val int
     Next *ListNode
}

func (n *ListNode) print() {
	cur :=n
	arr := []int{}
	for cur != nil {
		arr = append(arr,cur.Val)
		cur = cur.Next
	}

	fmt.Println(arr)
}

func buildList(arr []int) *ListNode {
	dum := &ListNode{}
	cur := dum
	for _,n := range arr {
		cur.Next = &ListNode{
			Val:n,
		}

		cur = cur.Next
	}

	return dum.Next
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	// 思路： prev 指向翻转的节点的前一个节点
	// 	翻转开始的节点为slow，终点为fast
	//  next 为fast的下一个节点，截断 slow之前和fast之后的指向
	// 	翻转 slow-fast 部分, 然后 prev.Next = fast
	// slow.Next = next
	// prev = slow
	// slow,fast = next,next
	dum := &ListNode{Next: head}
	prev := dum
	slow := head
	fast := head
	cnt := 0
	for fast != nil {
		if cnt < k-1 {
			cnt++
			if fast.Next != nil {
				fast = fast.Next
			} else {
				break
			}
			continue
		}

		// cnt == k-1
		next := fast.Next
		// 断开关系
		fast.Next = nil
		prev.Next = nil

		// 翻转
		// slow.print()
		reverse(slow)
		// fast.print()
		prev.Next = fast
		// prev.print()
		slow.Next = next
		// prev.print()
		// slow.print()
		prev = slow
		slow = next
		fast = next
		cnt = 0
		// dum.print()
	}

	// head.print()
	// 剩下的翻转
	// 剩余大于0
	// if cnt > 1 {
	// 	next := fast.Next
	// 	// 断开关系
	// 	fast.Next = nil
	// 	prev.Next = nil

	// 	// 翻转
	// 	reverse(slow)
	// 	prev.Next = fast
	// 	slow.Next = next
	// 	prev = slow
	// 	slow = next
	// 	fast = next
	// 	cnt = 0
	// }

	return dum.Next
}

func reverse(head *ListNode) {
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

