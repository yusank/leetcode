/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
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

func mergeKLists(lists []*ListNode) *ListNode {
	// 1. find a list start with min value
	// 2. insert others node into this list in order
	var (
		targetListIdx = -1
	)

	for i, head := range lists {
		if targetListIdx == -1 {
			targetListIdx = i
			continue
		}

		if head.Val < lists[targetListIdx].Val {
			targetListIdx = i
			continue
		}
	}

	if targetListIdx == -1 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var head = lists[targetListIdx]
	for head != nil {
		for i, n := range lists {
			if i == targetListIdx {
				continue
			}

		}
	}

}

// transfer `from` node to next of `to` node and return next of `from` node.
func transferNode(from, to *ListNode) *ListNode {
	fromNext := from.Next
	from.Next = nil // remove from old list

	// to
	toNext := to.Next
	to.Next = from
	from.Next = toNext

	return fromNext
}

// @lc code=end
