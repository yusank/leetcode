/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

	answer2(head)

	// 思路： 将链表转换为线性列表，然后根据下标重排
	// nodes := make([]*ListNode,0)

	// for node := head;node != nil ; node = node.Next {
	// 	nodes = append(nodes,node)
	// }

	// i,j := 0, len(nodes)-1
	// for i < j {
	// 	nodes[i].Next = nodes[j]
	// 	i++
	// 	if i == j {
	// 		break
	// 	}

	// 	nodes[j].Next = nodes[i]
	// 	j--
	// }

	// nodes[i].Next = nil
}

// 第二种思路
// 1. 将数组切分为左右半段
// 2. 将右半部分反转
// 3. 合并链表
func answer2(head *ListNode) {
	mid := findMid(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2= listReverse(l2)
	mergeList(l1,l2)
}

func findMid(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func listReverse(head *ListNode) *ListNode{
	var (
		prev *ListNode
		cur = head
	)

	for cur != nil {
		tempNext := cur.Next
		cur.Next=prev
		prev = cur
		cur = tempNext
	}

	return prev
}

func mergeList(l1,l2 *ListNode) {
	var l1temp,l2temp *ListNode
	for l1 != nil && l2 != nil {
		l1temp = l1.Next
		l2temp = l2.Next

		l1.Next = l2
		l1 = l1temp

		l2.Next = l1
		l2 = l2temp
	}
}
// @lc code=end

