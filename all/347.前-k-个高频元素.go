/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2}
	fmt.Println(topKFrequent(nums, 2))
}

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	var (
		countMap = make(map[int]int)
		temp     = make(map[int][]int)
		arr      []int
		result   []int
	)

	for _, n := range nums {
		countMap[n] += 1
	}

	// fmt.Println(countMap)

	// h := &IHeap{}
	// heap.Init(h)

	for n, v := range countMap {
		arr = append(arr, v)
		sli, ok := temp[v]
		if !ok {
			sli = make([]int, 0)
		}

		sli = append(sli, n)
		temp[v] = sli
	}

	// fmt.Println(temp)

	top := topK(arr, k)
	// fmt.Println(top)
	for _, n := range top {
		result = append(result, temp[n][0])
		temp[n] = temp[n][1:]
	}

	return result

}

func topK(nums []int, k int) []int {
	// fmt.Println(nums)
	size := len(nums)
	nums1 := make([]int, 0)

	buildHeap(nums, size)
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		// fmt.Println(i)
		nums1 = append(nums1, nums[0])
		nums[0], nums[i] = nums[i], nums[0]
		size--
		maxHeapify(nums, 0, size)
	}

	return nums1
}

func buildHeap(nums []int, size int) {
	for i := size / 2; i >= 0; i-- {
		maxHeapify(nums, i, size)
	}
}

func maxHeapify(nums []int, i, size int) {
	var (
		l       = i<<1 + 1
		r       = i<<1 + 2
		largest = i
	)

	if l < size && nums[l] > nums[largest] {
		largest = l
	}

	if r < size && nums[r] > nums[largest] {
		largest = r
	}

	if i != largest {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, size)
	}
}

// type IHeap [][2]int

// func (h IHeap) Len() int           { return len(h) }
// func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
// func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IHeap) Push(x interface{}) {
// 	*h = append(*h, x.([2]int))
// }

// func (h *IHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// @lc code=end
