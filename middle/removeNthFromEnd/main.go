package main

import "fmt"

/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

	示例：

	给定一个链表: 1->2->3->4->5, 和 n = 2.

	当删除了倒数第二个节点后，链表变为 1->2->3->5.
	说明：

	给定的 n 保证是有效的。

	进阶：

	你能尝试使用一趟扫描实现吗？

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	var cur = l
	fmt.Println("--------------")
	for cur != nil {
		fmt.Printf("val:%v ptr:%p |", cur.Val, cur)
		cur = cur.Next
	}
	fmt.Println("")
}

func main() {
	head := getSimpleList()
	head = removeNthFromEndWithDoublePtr(head, 1)
	head.print()
}

func getSimpleList() *ListNode {
	var head = &ListNode{
		Val: 1,
	}

	var cur = head
	for i := 2; i < 6; i++ {
		next := &ListNode{
			Val: i,
		}

		cur.Next = next
		cur = next
	}

	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	var (
		curNode   = head
		listSlice = make([]*ListNode, 0)
	)

	for curNode != nil {
		listSlice = append(listSlice, curNode)
		curNode = curNode.Next
	}

	length := len(listSlice)
	index := length - n
	if index < 0 {
		return head
	}
	if index == 0 {
		return head.Next
	}
	lastNode := listSlice[index-1]
	lastNode.print()

	if index == length-1 {
		lastNode.Next = nil
		return head
	}

	nextNode := listSlice[index+1]
	lastNode.Next = nextNode

	return head
}

// 双指针方案
func removeNthFromEndWithDoublePtr(head *ListNode, n int) *ListNode {
	dumNode := &ListNode{
		Val:  0,
		Next: head,
	}

	var (
		first, second = dumNode, dumNode
	)

	for i := 0; i < n; i++ {
		if first == nil {
			return head
		}
		first = first.Next
	}

	for first.Next != nil {
		second = second.Next
		first = first.Next
	}

	second.Next = second.Next.Next
	return dumNode.Next
}
