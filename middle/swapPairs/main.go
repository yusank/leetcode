/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (68.80%)
 * Likes:    769
 * Dislikes: 0
 * Total Accepted:    207.2K
 * Total Submissions: 301.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0
 *
 *
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

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	slice := []int{1, 2, 3, 4}
	head := &ListNode{
		Val: 0,
	}

	var cur = head
	for _, v := range slice {
		cur.Next = &ListNode{
			Val: v,
		}

		cur = cur.Next
	}

	fmt.Println(head)
	result := swapPairs(head.Next)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		cur    = head
		last   *ListNode
		next   = cur.Next
		result *ListNode
	)

	for next != nil {
		swap(cur,next)
		if last != nil {
			last.Next = next
		}

		if result == nil {
			result = next	
		}

		cur = cur.Next
		last = next.Next
		if cur == nil {
			break
		}
		next= cur.Next
	}

	return result
}


func swap(a,b *ListNode) {
	a.Next, b.Next = b.Next,a
}

// 1 - 2 - 3 - 4
// 2 - 1 - 3 - 4

// @lc code=end
