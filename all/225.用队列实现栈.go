/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */
package main

import "log"

func main() {
	ms := Constructor()
	ms.Push(1)
	ms.Push(2)
	ms.Push(3)
	log.Println(ms.Top())
	log.Println(ms.Pop())
	log.Println(ms.Top())
	log.Println(ms.Empty())
}

// @lc code=start
type MyStack struct {
	// val []int
	// 考虑做双向链表
	head *listNode
	tail *listNode
}

type listNode struct {
	val  int
	next *listNode
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.head == nil {
		this.tail = &listNode{val: x}
		this.head = this.tail
		return
	}

	this.tail.next = &listNode{val: x}
	this.tail = this.tail.next
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.head.val
	this.head = this.head.next
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.head.val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.head != nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
