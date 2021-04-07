/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack []int
	minStack []int
}

func min(a,b int) int {
	if a>b {
		return b
	}

	return a
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack : []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack,val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack,min(top,val))
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

