package main

func main() {
	
}

type ListNode struct {
	Val int
	Next *ListNode
}


func mergeTwoLists(l1,l2 *ListNode) *ListNode {
	var newList =new(ListNode)

	cur := newList
	for l1 != nil && l2 != nil {
		var nextNode *ListNode
		if l1.Val > l2.Val {
			nextNode = l2
			l2 = l2.Next
		} else {
			nextNode = l1
			l1 = l1.Next
		}

		cur.Next = nextNode
		cur = nextNode
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return newList.Next
}