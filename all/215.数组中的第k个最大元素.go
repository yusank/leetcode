/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	size := len(nums)
	build(nums,size)
	for i := len(nums)-1;i >= len(nums)-k+1;i-- {
		nums[0],nums[i] = nums[i],nums[0] // 不断把堆的根挪到后面
		size--
		maxHeapify(nums,0,size)
	}

	return nums[0]
}

func build(nums []int,size int) {
	for i := size/2;i >= 0;i-- {
		maxHeapify(nums,i,size)
	}
}

func maxHeapify(nums []int,i, size int) {
	l,r,large := i*2+1,i*2+2,i
	if l < size && nums[l] > nums[large] {
		large = l
	}

	if r < size && nums[r] > nums[large] {
		large =r
	}

	if large != i {
		nums[i], nums[large] = nums[large],nums[i]
		maxHeapify(nums,large,size)
	}
}
// @lc code=end

